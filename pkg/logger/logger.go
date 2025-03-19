package logger

import (
	"context"

	"go.uber.org/zap"
)

type contextKey string

const (
	requestIDKey contextKey = "request_id"
)

type Logger struct {
	l *zap.Logger
}

func New(ctx context.Context) (context.Context, *Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, nil, err
	}

	l := &Logger{l: logger}

	ctx = context.WithValue(ctx, requestIDKey, "")

	return ctx, l, nil
}

func GetLoggerFromContext(ctx context.Context) *Logger {
	if ctx == nil {
		zap.L().Warn("Context is nil, returning no-op logger")
		return &Logger{l: zap.NewNop()}
	}

	if logger, ok := ctx.Value(requestIDKey).(*Logger); ok {
		return logger
	}

	zap.L().Warn("Logger not found in context, returning no-op logger")
	return &Logger{l: zap.NewNop()}
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		fields = append(fields, zap.String("request_id", requestID))
	}

	l.l.Info(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		fields = append(fields, zap.String("request_id", requestID))
	}

	l.l.Error(msg, fields...)
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		fields = append(fields, zap.String("request_id", requestID))
	}

	l.l.Debug(msg, fields...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		fields = append(fields, zap.String("request_id", requestID))
	}

	l.l.Warn(msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		fields = append(fields, zap.String("request_id", requestID))
	}

	l.l.Fatal(msg, fields...)
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{l: l.l.With(fields...)}
}
