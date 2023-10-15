package loggers

import (
	log "github.com/sirupsen/logrus"
	"github.com/zcubbs/logwrapper/logger"
)

type LogrusLogger struct {
	*log.Entry
	baseLogger *log.Logger
}

func NewLogrusLogger(name string) logger.Logger {
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{})
	return &LogrusLogger{l.WithField("logger_name", name), l}
}

func (l *LogrusLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.Entry.Debugf(msg, keysAndValues...)
}

func (l *LogrusLogger) Info(msg string, keysAndValues ...interface{}) {
	l.Entry.Infof(msg, keysAndValues...)
}

func (l *LogrusLogger) Error(msg string, keysAndValues ...interface{}) {
	l.Entry.Errorf(msg, keysAndValues...)
}

func (l *LogrusLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.Entry.Fatalf(msg, keysAndValues...)
}

func (l *LogrusLogger) SetFormat(format string) {
	switch format {
	case logger.JSONFormat:
		l.baseLogger.SetFormatter(&log.JSONFormatter{})
	default:
		l.baseLogger.SetFormatter(&log.TextFormatter{})
	}
}
