package main

import (
	"log"
	"net/http"
	"nodenet/internal/config"
	"nodenet/internal/controllers"
	"nodenet/internal/logging"
	"nodenet/internal/routes"
	"nodenet/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize logger
	logger := logging.NewLogger("logs/node.log")

	// Initialize service with initial data
	service := services.NewNodeService(cfg.InitialData, logger)

	// Initialize controller
	controller := controllers.NewNodeController(service)

	// Set up routes
	routes.InitializeRoutes(controller)

	// Start server
	log.Println("Server starting on port:", cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
