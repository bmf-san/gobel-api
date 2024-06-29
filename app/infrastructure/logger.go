package infrastructure

import (
	"context"
	"os"

	"log/slog"
)

// Logger represents the singular of logger.
type Logger struct {
	*slog.Logger
}

// NewLogger creates a logger.
func NewLogger(level int) *Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(level),
	})
	logger := slog.New(handler)
	return &Logger{
		logger,
	}
}

func (l *Logger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.Logger.ErrorContext(ctx, msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.Logger.InfoContext(ctx, msg, args...)
}
