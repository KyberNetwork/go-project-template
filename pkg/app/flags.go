package app

import (
	"time"

	"github.com/urfave/cli"
)

const (
	timeoutFlag    = "timeout"
	defaultTimeout = 5 * time.Second
)

// NewTimeoutFlag return client flag to config timeout.
func NewTimeoutFlag() cli.Flag {
	return cli.DurationFlag{
		Name:   timeoutFlag,
		Usage:  "provide timeout for request",
		EnvVar: "TIMEOUT",
		Value:  defaultTimeout,
	}
}

// TimeoutFromContext return timeout from client configures.
func TimeoutFromContext(c *cli.Context) time.Duration {
	return c.Duration(timeoutFlag)
}
