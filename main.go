package main

import (
	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/kafka"
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
	log.Info("starting Hangout Storage Service", "logging-backend", cfg.Log.Backend)
	event, err := kafka.Consume(&cfg, log)
	if err != nil {
		log.Error("could not consume events from kakfa")
	}
	for e := range event {
		log.Info("file uploaded", "content-type", e.ContentType, "file name", e.Filename)
	}

}
