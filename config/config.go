package config

import (
	"log"
)

type Config struct {
	Kafka struct {
		Port    string `yaml:"port", envconfig:"KAFKA_PORT"`
		Host    string `yaml:"host", envconfig:"KAFKA_HOST"`
		Topic   string `yaml:"topic", envconfig:"KAFKA_TOPIC"`
		GroupId string `yaml:"group-id", envconfig:"KAFKA_GROUPID"`
	} `yaml:"kafka"`
	Log struct {
		Level   string `yaml:"level", envconfig:"LOG_LEVEL"`
		Backend string `yaml:"backend", envconfig:"LOG_BACKEND"`
	} `yaml:"log"`
	Path struct {
		Upload  string `yaml:"upload", envconfig:"PATH_UPLOAD"`
		Storage string `yaml:"storage", envconfig:"PATH_STORAGE"`
	} `yaml:"path"`
	Process struct {
		QueueLength  int `yaml:"queue-length", envconfig:"PROCESS_QUEUELENGTH"`
		PoolStrength int `yaml:"pool-strength", envconfig:"PROCESS_POOLSTRENGTH"`
	} `yaml:"process"`
}

// ? keeping this exception function here because when this function
// ? will execute loggers would not have been initialized
func configLoadError(source string, err *error) {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Print("Error in loading configuration from", source, " error: ", err)
	//os.Exit(1)
}
