package store

import "log/slog"

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Error(err error, msg string, kvs ...any) {
	slog.Error(msg, append(kvs, "err", err)...)
}
