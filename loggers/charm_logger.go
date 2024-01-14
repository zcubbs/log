package loggers

import (
	charmlog "github.com/charmbracelet/log"
	"github.com/zcubbs/log/structuredlogger"
	"os"
)

type CharmLogger struct {
	*charmlog.Logger
	baseLogger *charmlog.Logger
}

func NewCharmLogger(name string) structuredlogger.StructuredLogger {
	l := charmlog.New(os.Stderr)

	return &CharmLogger{
		Logger:     l.With("logger_name", name),
		baseLogger: l,
	}
}

func (c *CharmLogger) Debug(msg string, keysAndValues ...interface{}) {
	c.baseLogger.Debug(msg, keysAndValues...)
}

func (c *CharmLogger) Info(msg string, keysAndValues ...interface{}) {
	c.baseLogger.Info(msg, keysAndValues...)
}

func (c *CharmLogger) Warn(msg string, keysAndValues ...interface{}) {
	c.baseLogger.Warn(msg, keysAndValues...)
}

func (c *CharmLogger) Error(msg string, keysAndValues ...interface{}) {
	c.baseLogger.Error(msg, keysAndValues...)
}

func (c *CharmLogger) Fatal(msg string, keysAndValues ...interface{}) {
	c.baseLogger.Fatal(msg, keysAndValues...)
}

func (c *CharmLogger) SetFormat(format string) {
	switch format {
	case structuredlogger.JSONFormat:
		c.baseLogger.SetFormatter(charmlog.JSONFormatter)
	default:
		c.baseLogger.SetFormatter(charmlog.TextFormatter)
	}
}

func (c *CharmLogger) SetLevel(level string) {
	switch level {
	case structuredlogger.DebugLevel:
		c.baseLogger.SetLevel(charmlog.DebugLevel)
	case structuredlogger.InfoLevel:
		c.baseLogger.SetLevel(charmlog.InfoLevel)
	case structuredlogger.WarnLevel:
		c.baseLogger.SetLevel(charmlog.WarnLevel)
	case structuredlogger.ErrorLevel:
		c.baseLogger.SetLevel(charmlog.ErrorLevel)
	case structuredlogger.FatalLevel:
		c.baseLogger.SetLevel(charmlog.FatalLevel)
	}
}
