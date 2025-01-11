package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/files"
	"hangout.com/core/storage-service/kafka"
	"hangout.com/core/storage-service/logger"
	"hangout.com/core/storage-service/worker"
)

func main() {
	// Load configuration and initialize logger
	var cfg config.Config
	config.ReadFile(&cfg)
	config.ReadEnv(&cfg)
	log := logger.NewLogger(&cfg)
	// Create a base context with a cancel function for the application lifecycle
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancel is called on application exit

	// Handle OS signals for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Info("Received shutdown signal, cancelling context")
		cancel()
	}()

	log.Info("starting Hangout Storage Service", "logging-backend", cfg.Log.Backend)

	// Channel to handle incoming Kafka events
	eventChan := make(chan *files.File, cfg.Process.QueueLength)

	// Start the worker pool with the base context
	log.Info("Creating worker pool", "pool-strength", cfg.Process.PoolStrength)
	wp := worker.CreateWorkerPool(eventChan, ctx, &cfg, log)

	// Start the Kafka consumer
	log.Info("starting kafka consumer using ConsumerGroup API")
	err := kafka.StartConsumer(eventChan, ctx, &cfg, log)
	if err != nil {
		log.Error("Error starting Consumer Group")
	}

	// Wait for all workers to finish on shutdown
	wp.Wait()
	log.Info("Hangout Storage Service shut down gracefully")
}
