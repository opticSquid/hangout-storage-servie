package main

import (
	"fmt"

	"hangout.com/core/storage-service/config"
)

func main() {
	var cfg config.Config
	config.ReadFile(&cfg)
	config.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
}
