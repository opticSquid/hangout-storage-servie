package logger

import (
	"hangout.com/core/storage-service/config"
)

type Log interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

func NewLogger(cfg *config.Config) Log {
	if cfg.Log.Backend == "slog" {
		return NewSlogLogger(cfg)
	} else {
		return NewZeroLogger(cfg)
	}
}
