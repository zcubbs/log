package loggers

import (
	log "github.com/sirupsen/logrus"
	lw "github.com/zcubbs/logwrapper"
)

type LogrusLogger struct {
	*log.Entry
	baseLogger *log.Logger
}

func NewLogrusLogger(name string) lw.Logger {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	return &LogrusLogger{logger.WithField("logger_name", name), logger}
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
	case lw.JSONFormat:
		l.baseLogger.SetFormatter(&log.JSONFormatter{})
	default:
		l.baseLogger.SetFormatter(&log.TextFormatter{})
	}
}
