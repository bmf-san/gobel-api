package domain

import "context"

// A Logger is a logger interface.
type Logger interface {
	WithTraceID(context.Context) context.Context
	Error(string, ...any)
	ErrorContext(context.Context, string, ...any)
	Info(string, ...any)
	InfoContext(context.Context, string, ...any)
}
