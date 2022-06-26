package cli

import (
	"github.com/memclutter/goprotemplate/internal/cli/hooks"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:    "goprotemplate",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "dsnDb", Value: "postgresql://goprotemplate:goprotemplate@localhost:5432/goprotemplate?sslmode=disable", EnvVars: []string{"DSN_DB"}},
			&cli.BoolFlag{Name: "debug", Value: false, EnvVars: []string{"DEBUG"}},
			&cli.StringFlag{Name: "logLevel", Value: logrus.InfoLevel.String(), EnvVars: []string{"LOG_LEVEL"}},
		},
		Before:   hooks.Before,
		After:    hooks.After,
		Commands: cli.Commands{},
	}
}
