package app

import (
	"fmt"
	"io"
	"os"

	"github.com/KyberNetwork/cclog/lib/client"
	"github.com/TheZeroSlave/zapsentry"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	infoLevel  = "info"
	warnLevel  = "warn"
	errorLevel = "error"
	fatalLevel = "fatal"

	sentryDSNFlag      = "sentry-dsn"
	sentryLevelFlag    = "sentry-lv"
	defaultSentryLevel = errorLevel

	ccLogAddr = "cclog-addr"
	cclogName = "cclog-name"
)

// NewSentryFlags returns flags to init sentry client.
func NewSentryFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   sentryDSNFlag,
			EnvVar: "SENTRY_DSN",
			Usage:  "dsn for sentry client",
		},
		cli.StringFlag{
			Name:   sentryLevelFlag,
			EnvVar: "SENTRY_LEVEL",
			Usage:  "log level report message to sentry (info, error, warn, fatal)",
			Value:  defaultSentryLevel,
		},
		cli.StringFlag{
			Name:   ccLogAddr,
			Usage:  "cclog-address",
			Value:  "",
			EnvVar: "CCLOG_ADDR",
		},
		cli.StringFlag{
			Name:   cclogName,
			Usage:  "cclog-name",
			Value:  "sample-cclog-name",
			EnvVar: "CCLOG_NAME",
		},
	}
}

type syncer interface {
	Sync() error
}

// NewFlusher creates a new syncer from given syncer that log a error message if failed to sync.
func NewFlusher(s syncer) func() {
	return func() {
		// ignore the error as the sync function will always fail in Linux
		// https://github.com/uber-go/zap/issues/370
		_ = s.Sync()
	}
}

// NewLogger creates a new logger instance.
// The type of logger instance will be different with different application running modes.
func newLogger(c *cli.Context) (*zap.Logger, zap.AtomicLevel) {
	writers := []io.Writer{os.Stdout}
	logAddr := c.GlobalString(ccLogAddr)
	logName := c.GlobalString(cclogName)

	if logAddr != "" && logName != "" {
		ccw := client.NewAsyncLogClient(logName, logAddr, func(err error) {
			fmt.Fprintln(os.Stdout, "send log error", err)
		})

		writers = append(writers, &UnescapeWriter{w: ccw})
	}

	w := io.MultiWriter(writers...)
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.CallerKey = "caller"

	encoder := zapcore.NewConsoleEncoder(config)
	cc := zap.New(zapcore.NewCore(encoder, zapcore.AddSync(w), atom), zap.AddCaller())

	return cc, atom
}

// NewLogger creates a new sugared logger and a flush function. The flush function should be
// called by consumer before quitting application.
// This function should be use most of the time unless
// the application requires extensive performance, in this case use NewLogger.
func NewLogger(c *cli.Context) (*zap.Logger, zap.AtomicLevel, func(), error) {
	logger, atom := newLogger(c)

	// init sentry if flag dsn exists
	if len(c.String(sentryDSNFlag)) != 0 {
		sentryClient, err := sentry.NewClient(
			sentry.ClientOptions{
				Dsn: c.String(sentryDSNFlag),
			},
		)
		if err != nil {
			return nil, atom, nil, fmt.Errorf("failed to init sentry client: %w", err)
		}

		cfg := zapsentry.Configuration{
			DisableStacktrace: false,
		}

		switch c.String(sentryLevelFlag) {
		case infoLevel:
			cfg.Level = zapcore.InfoLevel
		case warnLevel:
			cfg.Level = zapcore.WarnLevel
		case errorLevel:
			cfg.Level = zapcore.ErrorLevel
		case fatalLevel:
			cfg.Level = zapcore.FatalLevel
		default:
			return nil, atom, nil, errors.Errorf("invalid log level %v", c.String(sentryLevelFlag))
		}

		core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromClient(sentryClient))
		if err != nil {
			return nil, atom, nil, fmt.Errorf("failed to init zap sentry: %w", err)
		}
		// attach to logger core
		logger = zapsentry.AttachCoreToLogger(core, logger)
	}

	return logger, atom, NewFlusher(logger), nil
}
