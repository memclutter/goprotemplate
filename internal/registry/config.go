package registry

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var Config Configuration

type Configuration struct {
	Debug    bool
	LogLevel logrus.Level
}

func (c *Configuration) SetLogLevel(ll string) error {
	lvl, err := logrus.ParseLevel(ll)
	if err != nil {
		return fmt.Errorf("set log level error: %v", err)
	}
	c.LogLevel = lvl
	return nil
}
