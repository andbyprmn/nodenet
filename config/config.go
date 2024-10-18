package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	//your config in here ...
}

func LoadConfig() Config {
	file, err := os.Open("your_config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

	return config
}
