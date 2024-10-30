package config

import (
	"log"
	"os"
)

type Config struct {
	Kafka struct {
		Port    string `yaml:"port", envconfig:"KAFKA_PORT"`
		Host    string `yaml:"host", envconfig:"KAFKA_HOST"`
		Topic   string `yaml:"topic", envconfig:"KAFKA_TOPIC"`
		GroupId string `yaml: "group-id", envconfig: "KAFKA_GROUP-ID"`
	} `yaml:"kafka"`
	Log struct {
		Level   string `yaml:"level", envconfig: "LOG_LEVEL"`
		Backend string `yaml: "backend", envconfig: LOG_BACKEND`
	} `yaml:"log"`
	Hangout struct {
		Media struct {
			UploadPath string `yaml: "upload-path", envconfig: "HANGOUT_MEDIA_UPLOAD-PATH"`
			QLength    int    `yaml: "queue-length", envconfig: "HANGOUT_MEDIA_QUEUE-LENGTH"`
		} `yaml: "media"`
	} `yaml: "hangout"`
}

// ? keeping this exception function here because when this function
// ? will execute loggers would not have been initialized
func configLoadError(err *error) {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Fatal("Error in loading configuration", err)
	os.Exit(1)
}
