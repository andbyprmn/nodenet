package main

import (
	"log"
	"net/http"
	"nodenet/config"
	"nodenet/handler"
	customLog "nodenet/log"
	"nodenet/model"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Set up logging
	customLog.SetupLogger(cfg.LogFile)

	// Create a new store
	store, err := model.NewStore(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Create a new router
	router := httprouter.New()

	// Create an API instance and register routes
	api := handler.NewAPI(store)
	api.RegisterRoutes(router)

	// Start the server
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
