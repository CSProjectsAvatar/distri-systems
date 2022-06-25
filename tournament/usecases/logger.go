package usecases

type Logger interface {
	Errorf(format string, args ...any)
	Infof(format string, args ...any)
	Info(msg string, fields LogArgs)
	Error(msg string, fields LogArgs)
}

type LogArgs map[string]any
