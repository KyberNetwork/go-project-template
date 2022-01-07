package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

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
	_, flush, err := libapp.NewSugaredLogger(c)
	if err != nil {
		return err
	}
	defer flush()
	// TODO: add logic here
	return nil
}
