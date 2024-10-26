package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

func ReadEnv(cfg *Config) {
	log.Println("Loading configuration from env variables")
	err := envconfig.Process("", cfg)
	if err != nil {
		configLoadError(&err)
	}
	log.Println("Configuration loading complete...")
}
