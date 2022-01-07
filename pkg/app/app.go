package app

import (
	"fmt"

	"github.com/urfave/cli"
)

const (
	modeFlag = "mode"

	developmentMode = "develop"
	productionMode  = "production"
)

var (
	validRunningModes = map[string]struct{}{
		developmentMode: {},
		productionMode:  {},
	}
)

// NewApp creates a new cli App instance with common flags pre-loaded.
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  modeFlag,
			Usage: "app running mode",
			Value: developmentMode,
		},
	}
	app.Flags = append(app.Flags, NewSentryFlags()...)
	return app
}

func modeFromContext(c *cli.Context) (string, error) {
	mode := c.GlobalString(modeFlag)
	_, ok := validRunningModes[mode]
	if !ok {
		return "", fmt.Errorf("invalid running mode: %q", c.GlobalString(modeFlag))
	}
	return mode, nil
}

// Validate validates common application configuration flags.
func Validate(c *cli.Context) error {
	_, err := modeFromContext(c)
	return err
}

// IsDevelopmentMode returns true if application is running in development mode.
func IsDevelopmentMode(c *cli.Context) (bool, error) {
	mode, err := modeFromContext(c)
	if err != nil {
		return false, err
	}
	return mode == developmentMode, nil
}

// IsProductionMode returns true if application is running in Production mode.
func IsProductionMode(c *cli.Context) (bool, error) {
	mode, err := modeFromContext(c)
	if err != nil {
		return false, err
	}
	return mode == productionMode, nil
}
