package loggers

import (
	lw "github.com/zcubbs/logwrapper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger(name string) lw.Logger {
	zapLogger := &ZapLogger{}
	zapLogger.initLogger(name, lw.TextFormat) // default to text format
	return zapLogger
}

func (z *ZapLogger) initLogger(name string, format string) {
	var logger *zap.Logger
	var err error

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string{"stdout"}

	switch format {
	case lw.JSONFormat:
		cfg.Encoding = "json"
	case lw.TextFormat:
		cfg.Encoding = "console"
	}

	logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	z.logger = logger.Sugar().Named(name)
}

func (l *ZapLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *ZapLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (z *ZapLogger) SetFormat(format string) {
	z.initLogger(z.logger.Desugar().Name(), format)
}
