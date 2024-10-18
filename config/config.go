package config

import (
	"encoding/json"
	"os"
)

// Config struct untuk menyimpan konfigurasi aplikasi
type Config struct {
	Port    string `json:"port"`
	LogFile string `json:"log_file"`
	DBHost  string `json:"db_host"`
	DBPort  string `json:"db_port"`
	DBUser  string `json:"db_user"`
	DBPass  string `json:"db_pass"`
	DBName  string `json:"db_name"`
}

// LoadConfig untuk memuat konfigurasi dari file JSON
func LoadConfig(filePath string) (Config, error) {
	var config Config
	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return config, err
	}
	return config, nil
}
