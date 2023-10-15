package main

import (
	"github.com/zcubbs/logwrapper"
	"github.com/zcubbs/logwrapper/logger"
)

func main() {
	// Creating a new Logrus logger in JSON format
	logrusLogger := logwrapper.NewLogger(logger.LogrusLoggerType, "LogrusLogger", logger.JSONFormat)
	logrusLogger.Info("This is a Logrus logger in JSON format")

	// Switching the Logrus logger to text format
	logrusLogger.SetFormat(logger.TextFormat)
	logrusLogger.Info("This is a Logrus logger in text format")

	// Creating a new Zap logger in JSON format
	zapLogger := logwrapper.NewLogger(logger.ZapLoggerType, "ZapLogger", logger.JSONFormat)
	zapLogger.Info("This is a Zap logger in JSON format")

	// Switching the Zap logger to text format
	zapLogger.SetFormat(logger.TextFormat)
	zapLogger.Info("This is a Zap logger in text format")

	// Creating a new Standard logger in JSON format
	stdLogger := logwrapper.NewLogger(logger.StdLoggerType, "StdLogger", logger.JSONFormat)
	stdLogger.Info("This is a Standard logger in JSON format")

	// Switching the Standard logger to text format
	stdLogger.SetFormat(logger.TextFormat)
	stdLogger.Info("This is a Standard logger in text format")
}
