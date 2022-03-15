package app

import (
	"github.com/urfave/cli"
)

// NewApp creates a new cli App instance with common flags pre-loaded.
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Flags = NewSentryFlags()

	return app
}
