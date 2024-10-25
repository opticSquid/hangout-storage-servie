package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"hangout.com/core/storage-service/exceptions"
)

func ReadFile(cfg *Config) {
	// file path relative to project root directory
	f, err := os.Open("./resources/application.yaml")
	if err != nil {
		exceptions.ProcessError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		exceptions.ProcessError(err)
	}
}
