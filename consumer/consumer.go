package consumer

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/exceptions"
	"hangout.com/core/storage-service/logger"
)

func Consume(cfg *config.Config, log logger.Log) {
	msgCount := 0
	worker, err := connectConsumer(cfg)
	if err != nil {
		exceptions.KafkaConnectError(&err, log)
	}
	consumer, err := worker.ConsumePartition(cfg.Kafka.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		exceptions.KafkaConsumerError(&err, log)
	}
	log.Info("Kafka consumer started")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	doneChan := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Error("consumer runtime error occoured", "error", err)
			case msg := <-consumer.Messages():
				msgCount++
				log.Debug("Event consumed", "order count", msgCount, "topic", msg.Topic, "event", msg.Value)
				// all values are byte slices convert them and handle here
			case <-sigChan:
				log.Warn("OS Interruption received")
				doneChan <- struct{}{}
			}
		}
	}()

	<-doneChan
	log.Debug("routine completed", "prcessed event count", msgCount)

	if err := worker.Close(); err != nil {
		exceptions.KafkaConnectError(&err, log)
	}
}

func connectConsumer(cfg *config.Config) (sarama.Consumer, error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Consumer.Return.Errors = true
	broker := []string{fmt.Sprintf("%s:%s", cfg.Kafka.Host, cfg.Kafka.Port)}
	return sarama.NewConsumer(broker, kafkaConfig)
}
