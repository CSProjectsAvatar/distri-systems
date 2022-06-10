package usecases

type Logger interface {
	Error(format string, args ...any)
	Info(format string, args ...any)
}
