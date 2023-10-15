package loggers

import (
	"github.com/charmbracelet/log"
	"github.com/zcubbs/logwrapper/logger"
	"os"
)

type CharmLogger struct {
	*log.Logger
	baseLogger *log.Logger
}

func NewCharmLogger(name string) logger.Logger {
	l := log.New(os.Stderr)

	return &CharmLogger{
		Logger:     l.With("logger_name", name),
		baseLogger: l,
	}
}

func (c *CharmLogger) Debug(msg string, keysAndValues ...interface{}) {
	log.Debug(msg, keysAndValues...)
}

func (c *CharmLogger) Info(msg string, keysAndValues ...interface{}) {
	log.Info(msg, keysAndValues...)
}

func (c *CharmLogger) Warn(msg string, keysAndValues ...interface{}) {
	log.Warn(msg, keysAndValues...)
}

func (c *CharmLogger) Error(msg string, keysAndValues ...interface{}) {
	log.Error(msg, keysAndValues...)
}

func (c *CharmLogger) Fatal(msg string, keysAndValues ...interface{}) {
	log.Fatal(msg, keysAndValues...)
}

func (c *CharmLogger) SetFormat(format string) {
	switch format {
	case logger.JSONFormat:
		c.baseLogger.SetFormatter(log.JSONFormatter)
	default:
		c.baseLogger.SetFormatter(log.TextFormatter)
	}
}

func (c *CharmLogger) SetLevel(level string) {
	switch level {
	case logger.DebugLevel:
		c.baseLogger.SetLevel(log.DebugLevel)
	case logger.InfoLevel:
		c.baseLogger.SetLevel(log.InfoLevel)
	case logger.WarnLevel:
		c.baseLogger.SetLevel(log.WarnLevel)
	case logger.ErrorLevel:
		c.baseLogger.SetLevel(log.ErrorLevel)
	case logger.FatalLevel:
		c.baseLogger.SetLevel(log.FatalLevel)
	}
}
