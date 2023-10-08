package infrastructure

import (
	"context"
	"os"

	"log/slog"

	"github.com/google/uuid"
)

// Logger represents the singular of logger.
type Logger struct {
	*slog.Logger
}

// NewLogger creates a logger.
func NewLogger(level int) *Logger {
	handler := TraceIDHandler{slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(level),
	})}
	logger := slog.New(handler)
	return &Logger{
		logger,
	}
}

// WithTraceID returns a context with a trace id.
func (l *Logger) WithTraceID(ctx context.Context) context.Context {
	uuid, _ := uuid.NewRandom()
	return context.WithValue(ctx, ctxTraceIDKey, uuid.String())
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

// TraceIDHandler represents the singular of trace id handler.
type TraceIDHandler struct {
	slog.Handler
}

type ctxTraceID struct{}

var ctxTraceIDKey = ctxTraceID{}

// Handle implements slog.Handler.
func (t TraceIDHandler) Handle(ctx context.Context, r slog.Record) error {
	tid, ok := ctx.Value(ctxTraceIDKey).(string)
	if ok {
		r.AddAttrs(slog.String("trace_id", tid))
	}
	return t.Handler.Handle(ctx, r)
}
