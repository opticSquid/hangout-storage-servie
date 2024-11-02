package files

import (
	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/logger"
)

type image struct {
	filename string
}

func (i *image) processMedia(cfg *config.Config, log logger.Log) error {
	return nil
}
