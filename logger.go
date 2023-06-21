package log

import (
	"context"
	"os"
)

// A Logger provides fast, leveled, structured logging
type Logger interface {
	Debug(context.Context, string, ...Field)
	Debugw(context.Context, string, ...interface{})
	Info(context.Context, string, ...Field)
	Infow(context.Context, string, ...interface{})
	Warn(context.Context, string, ...Field)
	Warnw(context.Context, string, ...interface{})
	Error(context.Context, string, ...Field)
	Errorw(context.Context, string, ...interface{})
	Panic(context.Context, string, ...Field)
	Panicw(context.Context, string, ...interface{})
	Fatal(context.Context, string, ...Field)
	Fatalw(context.Context, string, ...interface{})
}

var (
	globalLogger Logger = New()
)

func New(options ...Option) Logger {
	parameter := &Parameter{
		Writer:   os.Stdout,
		LogLevel: LevelDebug,
	}

	for _, option := range options {
		option(parameter)
	}

	return NewZapLogger(parameter)
}

// ReplaceGlobals replace package level logger
func ReplaceGlobals(logger Logger) {
	globalLogger = logger
}

func Debug(ctx context.Context, message string, fields ...Field) {
	globalLogger.Debug(ctx, message, fields...)
}

func Debugw(ctx context.Context, message string, keyAndValues ...interface{}) {
	globalLogger.Debugw(ctx, message, keyAndValues...)
}

func Info(ctx context.Context, message string, fields ...Field) {
	globalLogger.Info(ctx, message, fields...)
}

func Infow(ctx context.Context, message string, keyAndValues ...interface{}) {
	globalLogger.Infow(ctx, message, keyAndValues...)
}

func Warn(ctx context.Context, message string, fields ...Field) {
	globalLogger.Warn(ctx, message, fields...)
}

func Warnw(ctx context.Context, message string, keyAndValues ...interface{}) {
	globalLogger.Warnw(ctx, message, keyAndValues...)
}

func Error(ctx context.Context, message string, fields ...Field) {
	globalLogger.Error(ctx, message, fields...)
}

func Errorw(ctx context.Context, message string, keyAndValues ...interface{}) {
	globalLogger.Errorw(ctx, message, keyAndValues...)
}

func Panic(ctx context.Context, message string, fields ...Field) {
	globalLogger.Panic(ctx, message, fields...)
}

func Panicw(ctx context.Context, message string, keyAndValues ...interface{}) {
	globalLogger.Panicw(ctx, message, keyAndValues...)
}

func Fatal(ctx context.Context, message string, fields ...Field) {
	globalLogger.Fatal(ctx, message, fields...)
}

func Fatalw(ctx context.Context, message string, keyAndValues ...interface{}) {
	globalLogger.Fatalw(ctx, message, keyAndValues...)
}
