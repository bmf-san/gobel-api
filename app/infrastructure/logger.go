package infrastructure

import (
	"github.com/bmf-san/golem"
)

// A Logger represents a logger.
type Logger struct {
	logger *golem.Logger
}

// NewLogger creates a logger.
func NewLogger() *Logger {
	return &Logger{
		logger: golem.NewLogger(),
	}
}

// Fatal outputs a fatal level log.
func (l *Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

// Error outputs a fatal level log.
func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

// Warn outputs a fatal level log.
func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

// Info outputs a fatal level log.
func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}
