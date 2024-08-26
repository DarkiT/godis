package logger

import (
	"github.com/darkit/slog"
)

var log *slog.Logger

func init() {
	log = slog.Default("Godis")
}

// Debug 记录Debug级别的日志。
func Debug(msg string, args ...any) {
	log.Debug(msg, args...)
}

// Info 记录Info级别的日志。
func Info(msg string, args ...any) {
	log.Info(msg, args...)
}

// Warn 记录Warn级别的日志。
func Warn(msg string, args ...any) {
	log.Warn(msg, args...)
}

// Warn 记录Warn级别的日志。
func Warnf(msg string, args ...any) {
	log.Warnf(msg, args...)
}

// Error 记录Error级别的日志。
func Error(msg string, args ...any) {
	log.Error(msg, args...)
}

// Fatal 记录Fatal级别的日志，并退出程序。
func Fatal(msg string, args ...any) {
	log.Fatal(msg, args...)
}

// Debugf 记录格式化的Debug级别的日志。
func Debugf(format string, args ...any) {
	log.Debugf(format, args...)
}

// Infof 记录格式化的Info级别的日志。
func Infof(format string, args ...any) {
	log.Infof(format, args...)
}

// Errorf 记录格式化的Error级别的日志。
func Errorf(format string, args ...any) {
	log.Errorf(format, args...)
}
