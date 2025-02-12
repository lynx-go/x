package log

import (
	"context"
	"log/slog"
	"runtime"
	"time"
)

type logCtx struct {
}

var logKey = logCtx{}

func WithContext(ctx context.Context, kwargs ...interface{}) context.Context {
	return Context(ctx, FromContext(ctx), kwargs...)
}

func Context(ctx context.Context, logger *slog.Logger, kwargs ...interface{}) context.Context {
	logger = logger.With(kwargs...)
	return context.WithValue(ctx, logKey, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(logKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}

func DebugContext(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelDebug, msg, args...)
}

func log(ctx context.Context, level slog.Level, msg string, args ...any) {

	logger := FromContext(ctx)
	if !logger.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	runtime.Callers(3, pcs[:])
	pc = pcs[0]
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = logger.Handler().Handle(ctx, r)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelInfo, msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelWarn, msg, args...)
}

func ErrorContext(ctx context.Context, msg string, err error, args ...any) {
	newArgs := []any{}
	newArgs = append(newArgs, args...)
	if err != nil {
		newArgs = []any{"error", err.Error()}
	}
	log(ctx, slog.LevelError, msg, newArgs...)
}
