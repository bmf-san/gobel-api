package domain

import "context"

// A Logger is a logger interface.
type Logger interface {
	Error(string, ...any)
	ErrorContext(context.Context, string, ...any)
	Info(string, ...any)
	InfoContext(context.Context, string, ...any)
}
