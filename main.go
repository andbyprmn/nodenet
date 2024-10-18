package main

import (
	"log"
	"net/http"
	"nodenet/config"
	"nodenet/handler"
	customLog "nodenet/log"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Load configuration
	conf := config.LoadConfig()

	// Initialize logger
	err := customLog.InitLogger(conf.LogFile)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer customLog.CloseLogger()

	// Initialize HTTP router
	router := httprouter.New()
	router.GET("/api/v1/keys/:key", handler.GetKey)
	router.POST("/api/v1/keys", handler.SetKey)

	// Start HTTP server
	log.Printf("Starting server on port %s...", conf.Port)
	log.Fatal(http.ListenAndServe(":"+conf.Port, router))
}
