package config

type Config struct {
	Kafka struct {
		Port string `yaml:"port", envconfig:"KAFKA_PORT"`
		Host string `yaml:"host", envconfig:"KAFKA_HOST"`
	} `yaml:"kafka"`
}
