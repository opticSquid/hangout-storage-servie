package config

import (
	"github.com/kelseyhightower/envconfig"
	"hangout.com/core/storage-service/exceptions"
)

func ReadEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		exceptions.ProcessError(err)
	}
}
