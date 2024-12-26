package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFile(cfg *Config) {
	log.SetFlags(log.Ldate | log.Lshortfile)
	// file path relative to project root directory
	f, err := os.Open("./resources/application.yaml")
	log.Println("Loading configuration from file")
	if err != nil {
		configLoadError("file", &err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		configLoadError("file", &err)
	}
	log.Println("Configuration loading complete from file...")
}
