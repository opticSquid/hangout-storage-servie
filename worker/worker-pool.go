package worker

import (
	"context"
	"sync"
	"time"

	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/files"
	"hangout.com/core/storage-service/logger"
)

type WorkerPool struct {
	eventChan <-chan *files.File
	wg        *sync.WaitGroup
	ctx       context.Context
}

func CreateWorkerPool(eventChan <-chan *files.File, ctx context.Context, cfg config.Config, log logger.Log) *WorkerPool {
	wp := &WorkerPool{eventChan: eventChan, wg: &sync.WaitGroup{}, ctx: ctx}
	for i := 0; i < cfg.Hangout.WorkerPool.Strength; i++ {
		log.Debug("spawning worker", "worker-id", i)
		wp.wg.Add(1)
		go wp.worker(i, log)
	}
	return wp
}

func (wp *WorkerPool) worker(workerId int, log logger.Log) {
	defer wp.wg.Done()
	for {
		select {
		case event, ok := <-wp.eventChan:
			if !ok {
				log.Info("Event channel closed, stopping worker", "worker-id", workerId)
				return
			}
			log.Info("starting file processing", "file-name", event.Filename, "worker-id", workerId)
			time.Sleep(2 * time.Second)
			log.Info("finished file processing", "file-name", event.Filename, "worker-id", workerId)
		case <-wp.ctx.Done():
			log.Info("Context cancelled, stopping worker", "worker-id", workerId)
			return
		}
	}
}

// Wait ensures all workers complete processing before the program exits
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}
