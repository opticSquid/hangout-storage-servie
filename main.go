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
	log := logger.NewLogger(&cfg)
	log.Info("starting Hangout Storage Service", "logging-backend", cfg.Log.Backend)
	log.Info("printing current configurtion", "config", cfg)
	log.Info("starting kafka consumer using ConsumerGroup API")
	files, err := kafka.StartConsumer(&cfg, log)
	if err != nil {
		log.Error("Error starting Consumer Group")
	}
	for file := range files {
		log.Info("file uploaded", "filename", file.Filename)
	}
}
