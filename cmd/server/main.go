package main

import (
	"log"

	"github.com/TFMV/gogeo/internal/handlers"
	services "github.com/TFMV/gogeo/internal/services"
	"github.com/TFMV/gogeo/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Initialize Google Maps client
	mapsClient, err := services.NewMapsClient(cfg.GoogleMapsAPIKey)
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
