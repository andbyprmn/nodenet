package main

import (
	"log"
	"nodenet/internal/api"
	"nodenet/internal/config"
	"nodenet/internal/data"
	"nodenet/internal/logging"
)

func main() {
	// Memuat file konfigurasi
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Inisialisasi penyimpanan key-value dan logger
	store := data.NewKeyValueStore(cfg.InitialData)
	logger := logging.NewLogger("logs/node.log")

	// Menjalankan server API
	apiServer := api.NewAPIServer(store, logger)
	apiServer.StartServer(cfg.Port)
}
