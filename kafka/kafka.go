package kafka

import (
	"context"
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

// This implements Sarama consumer group API
// Supports multi instance. All instances will join same consumer group.
// Single consumer per instance
func StartConsumer(cfg *config.Config, log logger.Log) (chan *media.Media, error) {
	log.Debug("starting to configure kafka client")
	consumerGroup, err := configureKafka(cfg)
	if err != nil {
		exceptions.KafkaConnectError("could not setup kafka connection", &err, log)
	}
	log.Info("kafka consumerGroup configured")
	eventChan := make(chan *media.Media, cfg.Hangout.Media.QLength)
	go consume(eventChan, consumerGroup, cfg, log)
	return eventChan, err
}
func configureKafka(cfg *config.Config) (sarama.ConsumerGroup, error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Version = sarama.DefaultVersion
	kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	brokers := []string{fmt.Sprintf("%s:%s", cfg.Kafka.Host, cfg.Kafka.Port)}
	return sarama.NewConsumerGroup(brokers, cfg.Kafka.GroupId, kafkaConfig)
}

func consume(eventChan chan *media.Media, consumerGroup sarama.ConsumerGroup, cfg *config.Config, log logger.Log) {
	defer close(eventChan)
	defer consumerGroup.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS interrupts for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()
	// this is the main function that calls the consumer
	handler := &ConsumerGroupHandler{Files: eventChan, log: log}
	for {
		if err := consumerGroup.Consume(ctx, []string{cfg.Kafka.Topic}, handler); err != nil {
			exceptions.KafkaConsumerError("Error in consumer loop", &err, log)
			return
		}
		if ctx.Err() != nil {
			log.Info("Context cancelled, stopping consumer")
			return
		}
	}
}
