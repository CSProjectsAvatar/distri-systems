package domain

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Info
	Error
)

type Logger interface {
	Trace(msg string, fields LogArgs)
	Debug(msg string, fields LogArgs)

	// Deprecated: It's better to assign names to args. Use Info instead.
	Infof(format string, args ...any)
	Info(msg string, fields LogArgs)

	// Deprecated: It's better to assign names to args. Use Error instead.
	Errorf(format string, args ...any)
	Error(msg string, fields LogArgs)

	// ToFile returns a logger which logs to a file. This is a fluent API.
	ToFile() Logger

	// WithLevel returns a logger with the specified level. This is a fluent API.
	// Default level is Info.
	WithLevel(level LogLevel) Logger
}

type LogArgs map[string]any
