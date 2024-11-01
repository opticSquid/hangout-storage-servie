package files

import (
	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/logger"
)

type media interface {
	processMedia(cfg *config.Config, log logger.Log)
}
