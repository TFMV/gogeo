package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tfmv/gogeo/internal/handlers"
	"github.com/tfmv/gogeo/pkg/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Initialize Google Maps client
	mapsClient, err := handlers.NewMapsClient(cfg.GoogleMapsAPIKey)
	if err != nil {
		log.Fatalf("Could not create maps client: %v", err)
	}

	r := gin.Default()

	r.GET("/geocode", handlers.GeocodeHandler(mapsClient))
	r.GET("/reverse_geocode", handlers.ReverseGeocodeHandler(mapsClient))

	if err := r.Run(); err != nil {
		log.Fatalf("Server run failed: %v", err)
	}
}
