package config

import (
	"log"
	"os"
)

type Config struct {
	Kafka struct {
		Port  string `yaml:"port", envconfig:"KAFKA_PORT"`
		Host  string `yaml:"host", envconfig:"KAFKA_HOST"`
		Topic string `yaml:"topic", envconfig:"KAFKA_TOPIC"`
	} `yaml:"kafka"`
	Log struct {
		Level   string `yaml:"level", envconfig: "LOG_LEVEL"`
		Backend string `yaml: "backend", envconfig: LOG_BACKEND`
	} `yaml:"log"`
}

// ? keeping this exception function here because when this function
// ? will execute loggers would not have been initialized
func configLoadError(err *error) {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Fatal("Error in loading configuration", err)
	os.Exit(1)
}
