package kafka

import (
	"github.com/IBM/sarama"
	"hangout.com/core/storage-service/logger"
	"hangout.com/core/storage-service/media"
)

// ConsumerGroupHandler implements sarama.ConsumerGroupHandler
type ConsumerGroupHandler struct {
	Files chan<- *media.Media
	log   logger.Log
}

// Setup runs at the beginning of a new session, before ConsumeClaim
func (cgh *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup runs at the end of a session, once all ConsumeClaim goroutines have exited
func (cgh *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim starts a consumer loop for each partition assigned to this handler
func (cgh *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		event := media.Media{ContentType: string(message.Key), Filename: string(message.Value)}
		cgh.log.Debug("event consumed", "Topic", message.Topic, "Partition", message.Partition, "Offset", message.Offset, "Key", event.ContentType, "Value", event.Filename)
		cgh.Files <- &event
		session.MarkMessage(message, "")
	}
	return nil
}
