package exceptions

import (
	"os"

	"hangout.com/core/storage-service/logger"
)

func ProcessError(err *error, log logger.Log) {
	log.Error("Processing error occoured", "error", err)
	os.Exit(2)
}

func KafkaConnectError(err *error, log logger.Log) {
	log.Error("could not setup kafka connection", "error", err)
	os.Exit(3)
}

func KafkaConsumerError(err *error, log logger.Log) {
	log.Error("could not consume partition", "error", err)
	os.Exit(4)
}
