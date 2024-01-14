package loggers

import (
	charmlog "github.com/charmbracelet/log"
	"github.com/zcubbs/log/structuredlogger"
	"os"
)

type CharmLogger struct {
	*charmlog.Logger
}

func NewCharmLogger(_ string) structuredlogger.StructuredLogger {
	l := charmlog.NewWithOptions(os.Stderr, charmlog.Options{
		ReportTimestamp: true,
	})

	return &CharmLogger{
		Logger: l,
	}
}

func (c *CharmLogger) Debug(msg string, keysAndValues ...interface{}) {
	c.Logger.Debug(msg, keysAndValues...)
}

func (c *CharmLogger) Info(msg string, keysAndValues ...interface{}) {
	c.Logger.Info(msg, keysAndValues...)
}

func (c *CharmLogger) Warn(msg string, keysAndValues ...interface{}) {
	c.Logger.Warn(msg, keysAndValues...)
}

func (c *CharmLogger) Error(msg string, keysAndValues ...interface{}) {
	c.Logger.Error(msg, keysAndValues...)
}

func (c *CharmLogger) Fatal(msg string, keysAndValues ...interface{}) {
	c.Logger.Fatal(msg, keysAndValues...)
}

func (c *CharmLogger) SetFormat(format string) {
	switch format {
	case structuredlogger.JSONFormat:
		c.Logger.SetFormatter(charmlog.JSONFormatter)
	default:
		c.Logger.SetFormatter(charmlog.TextFormatter)
	}
}

func (c *CharmLogger) SetLevel(level string) {
	switch level {
	case structuredlogger.DebugLevel:
		c.Logger.SetLevel(charmlog.DebugLevel)
	case structuredlogger.InfoLevel:
		c.Logger.SetLevel(charmlog.InfoLevel)
	case structuredlogger.WarnLevel:
		c.Logger.SetLevel(charmlog.WarnLevel)
	case structuredlogger.ErrorLevel:
		c.Logger.SetLevel(charmlog.ErrorLevel)
	case structuredlogger.FatalLevel:
		c.Logger.SetLevel(charmlog.FatalLevel)
	}
}
