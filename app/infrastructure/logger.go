package infrastructure

import (
	"time"

	"github.com/bmf-san/golem"
)

// A Logger represents a logger.
type Logger struct {
	logger *golem.Logger
}

// NewLogger creates a logger.
func NewLogger(threshold int, location *time.Location) *Logger {
	return &Logger{
		logger: golem.NewLogger(threshold, location),
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
