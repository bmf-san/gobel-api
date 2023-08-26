package infrastructure

import (
	"os"

	"log/slog"
)

// NewLogger creates a logger.
func NewLogger(level int) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(level),
	})
	logger := slog.New(handler)
	return logger
}
