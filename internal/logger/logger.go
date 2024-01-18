package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

func Init(core zapcore.Core, options ...zap.Option) {
	globalLogger = zap.New(core, options...)
}

func Debug(msg string, fields ...zapcore.Field) {
	globalLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	globalLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	globalLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	globalLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	globalLogger.Fatal(msg, fields...)
}

func WithOptions(opts ...zap.Option) *zap.Logger {
	return globalLogger.WithOptions(opts...)
}
