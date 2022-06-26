package hooks

import (
	"database/sql"
	"fmt"
	"github.com/memclutter/goprotemplate/internal/models"
	"github.com/memclutter/goprotemplate/internal/registry"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/urfave/cli/v2"
)

func Before(c *cli.Context) error {
	registry.Config.Debug = c.Bool("debug")
	if err := registry.Config.SetLogLevel(c.String("logLevel")); err != nil {
		return fmt.Errorf("cli hooks before error: %v", err)
	}
	logrus.SetLevel(registry.Config.LogLevel)
	models.DB = bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.String("dsnDb")))), pgdialect.New())

	// For debug mode, enable sql queries log and debug log level
	if registry.Config.Debug {
		models.DB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
		if registry.Config.LogLevel < logrus.DebugLevel {
			logrus.SetLevel(logrus.DebugLevel)
		}
	}
	return nil
}
