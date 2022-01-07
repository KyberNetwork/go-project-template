package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"go.uber.org/zap"

	libapp "github.com/KyberNetwork/go-project-template/pkg/app"
)

func main() {
	app := libapp.NewApp()
	app.Name = "alert service go script runner"
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}

func run(c *cli.Context) error {
	logger, _, flush, err := libapp.NewLogger(c)
	if err != nil {
		return err
	}
	defer flush()
	zap.ReplaceGlobals(logger)
	l := logger.Sugar()
	l.Infow("app starting ..")
	// TODO: add logic here
	return nil
}
