package structuredlogger

// Logger Types
const (
	StdLoggerType    = "std"
	LogrusLoggerType = "logrus"
	ZapLoggerType    = "zap"
	CharmLoggerType  = "charm"
)

// Logger Formats
const (
	TextFormat = "text"
	JSONFormat = "json"
)

// Logger Levels
const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	FatalLevel = "fatal"
)
