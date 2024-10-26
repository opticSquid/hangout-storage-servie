package main

import (
	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/logger"
)

func main() {
	var cfg config.Config
	config.ReadFile(&cfg)
	config.ReadEnv(&cfg)
	var log logger.Log
	if cfg.Log.Backend == "slog" {
		log = logger.NewSlogLogger(&cfg)
	} else {
		log = logger.NewZeroLogger(&cfg)
	}
	log.Debug("logging is working", "log-backend", cfg.Log.Backend)
}
