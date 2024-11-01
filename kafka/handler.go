package kafka

import (
	"github.com/IBM/sarama"
	"hangout.com/core/storage-service/files"
	"hangout.com/core/storage-service/logger"
)

// ConsumerGroupHandler implements sarama.ConsumerGroupHandler
type ConsumerGroupHandler struct {
	Files chan<- *files.File
	log   logger.Log
}

// Setup runs at the beginning of a new session, before ConsumeClaim
func (cgh *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	cgh.log.Info("Consumer group session setup completed")
	return nil
}

// Cleanup runs at the end of a session, once all ConsumeClaim goroutines have exited
func (cgh *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	cgh.log.Info("Consumer group session cleanup completed")
	return nil
}

// ConsumeClaim starts a consumer loop for each partition assigned to this handler
func (cgh *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		event := files.File{ContentType: string(message.Key), Filename: string(message.Value)}
		cgh.log.Debug("Event consumed",
			"Topic", message.Topic,
			"Partition", message.Partition,
			"Offset", message.Offset,
			"Key", event.ContentType,
			"Value", event.Filename,
		)
		// Send event to Files channel with non-blocking check
		select {
		case cgh.Files <- &event:
			session.MarkMessage(message, "")
		default:
			// Log a warning if the channel is full or has no active consumers
			cgh.log.Warn("File channel is full, unable to process event",
				"FileName", event.Filename,
				"ContentType", event.ContentType,
				"Partition", message.Partition,
				"Offset", message.Offset,
			)
		}
	}
	return nil
}
