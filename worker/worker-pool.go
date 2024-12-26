package worker

import (
	"context"
	"sync"

	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/files"
	"hangout.com/core/storage-service/logger"
)

type WorkerPool struct {
	eventChan <-chan *files.File
	wg        *sync.WaitGroup
	ctx       context.Context
	cfg       *config.Config
	log       logger.Log
}

func CreateWorkerPool(eventChan <-chan *files.File, ctx context.Context, cfg *config.Config, log logger.Log) *WorkerPool {
	wp := &WorkerPool{eventChan: eventChan, wg: &sync.WaitGroup{}, ctx: ctx, cfg: cfg, log: log}
	for i := 0; i < cfg.Process.PoolStrength; i++ {
		log.Debug("spawning worker", "worker-id", i)
		wp.wg.Add(1)
		go wp.worker(i)
	}
	return wp
}

func (wp *WorkerPool) worker(workerId int) {
	defer wp.wg.Done()
	for {
		select {
		case event, ok := <-wp.eventChan:
			if !ok {
				wp.log.Info("Event channel closed, stopping worker", "worker-id", workerId)
				return
			}
			wp.log.Info("starting file processing", "file-name", event.Filename, "user-id", event.UserId, "worker-id", workerId)
			// call file processing function here
			err := event.Process(wp.cfg, wp.log)
			if err != nil {
				wp.log.Error("could not process file", "error", err.Error())
			}
			wp.log.Info("finished file processing", "file-name", event.Filename, "worker-id", workerId)
		case <-wp.ctx.Done():
			wp.log.Info("Context cancelled, stopping worker", "worker-id", workerId)
			return
		}
	}
}

// Wait ensures all workers complete processing before the program exits
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}
