package domain

type Logger interface {
	Errorf(format string, args ...any)
	Infof(format string, args ...any)
	Info(msg string, fields LogArgs)
	Error(msg string, fields LogArgs)

	// ToFile returns a logger which logs to a file. This is a fluent API.
	ToFile() Logger
}

type LogArgs map[string]any
