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

func Context(ctx context.Context, logger *slog.Logger, kwargs ...interface{}) context.Context {
	logger = logger.With(kwargs...)
	return context.WithValue(ctx, logKey, logger)
}

func FromContext(ctx context.Context, kwargs ...interface{}) *slog.Logger {
	logger, ok := ctx.Value(logKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger.With(kwargs...)
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
	errmsg := "<nil>"
	if err != nil {
		errmsg = err.Error()
	}
	args = append(args, "error", errmsg)
	log(ctx, slog.LevelError, msg, args...)
}
