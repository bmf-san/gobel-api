package domain

// A Logger represents a logger.
type Logger interface {
	Fatal(msg string)
	Error(msg string)
	Warn(msg string)
	Info(msg string)
}
