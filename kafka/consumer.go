package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/exceptions"
	"hangout.com/core/storage-service/logger"
	"hangout.com/core/storage-service/media"
)

func Consume(cfg *config.Config, log logger.Log) (chan media.Media, error) {
	msgCount := 0
	worker, err := connectConsumer(cfg)
	if err != nil {
		exceptions.KafkaConnectError(&err, log)
	}
	consumer, err := worker.ConsumePartition(cfg.Kafka.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		exceptions.KafkaConsumerError(&err, log)
	}
	log.Info("Kafka consumer started")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	eventChan := make(chan media.Media, cfg.Hangout.Media.QLength)
	doneChan := make(chan struct{})
	go consumeEvents(consumer, eventChan, sigChan, doneChan, log, &msgCount)
	// Handle cleanup in a separate goroutine to keep function non-blocking
	go func() {
		<-doneChan
		log.Debug("Routine completed", "processed event count", msgCount)
		if err := worker.Close(); err != nil {
			exceptions.KafkaConnectError(&err, log)
		}
		close(eventChan) // Close eventChan when done
	}()
	return eventChan, err
}

func connectConsumer(cfg *config.Config) (sarama.Consumer, error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Consumer.Return.Errors = true
	broker := []string{fmt.Sprintf("%s:%s", cfg.Kafka.Host, cfg.Kafka.Port)}
	return sarama.NewConsumer(broker, kafkaConfig)
}

func consumeEvents(consumer sarama.PartitionConsumer, eventChan chan<- media.Media, sigChan <-chan os.Signal, doneChan chan<- struct{}, log logger.Log, msgCount *int) {
	for {
		select {
		case err := <-consumer.Errors():
			log.Error("Consumer runtime error occurred", "error", err)
		case msg := <-consumer.Messages():
			*msgCount++
			event := media.Media{ContentType: string(msg.Key), Filename: string(msg.Value)}
			log.Debug("Event consumed", "order count", *msgCount, "topic", msg.Topic)
			eventChan <- event
		case <-sigChan:
			log.Warn("OS Interruption received")
			doneChan <- struct{}{}
			return
		}
	}
}
