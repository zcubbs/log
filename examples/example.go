package main

import (
	"github.com/zcubbs/logwrapper"
)

func main() {
	// Creating a new Logrus logger in JSON format
	logrusLogger := logwrapper.NewLogger(logwrapper.LogrusLoggerType, "LogrusLogger", logwrapper.JSONFormat)
	logrusLogger.Info("This is a Logrus logger in JSON format")

	// Switching the Logrus logger to text format
	logrusLogger.SetFormat(logwrapper.TextFormat)
	logrusLogger.Info("This is a Logrus logger in text format")

	// Creating a new Zap logger in JSON format
	zapLogger := logwrapper.NewLogger(logwrapper.ZapLoggerType, "ZapLogger", logwrapper.JSONFormat)
	zapLogger.Info("This is a Zap logger in JSON format")

	// Switching the Zap logger to text format
	zapLogger.SetFormat(logwrapper.TextFormat)
	zapLogger.Info("This is a Zap logger in text format")

	// Creating a new Standard logger in JSON format
	stdLogger := logwrapper.NewLogger(logwrapper.StdLoggerType, "StdLogger", logwrapper.JSONFormat)
	stdLogger.Info("This is a Standard logger in JSON format")

	// Switching the Standard logger to text format
	stdLogger.SetFormat(logwrapper.TextFormat)
	stdLogger.Info("This is a Standard logger in text format")
}
