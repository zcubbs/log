package loggers

import (
	"github.com/charmbracelet/log"
	lw "github.com/zcubbs/logwrapper"
	"os"
)

type CharmLogger struct {
	*log.Logger
	baseLogger *log.Logger
}

func NewCharmLogger(name string) lw.Logger {
	logger := log.New(os.Stderr)

	return &CharmLogger{
		Logger:     logger.With("logger_name", name),
		baseLogger: logger,
	}
}

func (c *CharmLogger) Debug(msg string, keysAndValues ...interface{}) {
	log.Debug(msg, keysAndValues...)
}

func (c *CharmLogger) Info(msg string, keysAndValues ...interface{}) {
	log.Info(msg, keysAndValues...)
}

func (c *CharmLogger) Error(msg string, keysAndValues ...interface{}) {
	log.Error(msg, keysAndValues...)
}

func (c *CharmLogger) Fatal(msg string, keysAndValues ...interface{}) {
	log.Fatal(msg, keysAndValues...)
}

func (c *CharmLogger) SetFormat(format string) {
	switch format {
	case lw.JSONFormat:
		c.baseLogger.SetFormatter(log.JSONFormatter)
	default:
		c.baseLogger.SetFormatter(log.TextFormatter)
	}
}
