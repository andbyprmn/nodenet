package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port        string            `json:"port"`
	InitialData map[string]string `json:"initial_data"`
}

// LoadConfig memuat konfigurasi dari file JSON yang diberikan.
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return &config, nil
}
