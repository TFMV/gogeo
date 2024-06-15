package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"googlemaps.github.io/maps"
)

func GeocodeHandler(client *maps.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Query("address")
		if address == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Address is required"})
			return
		}

		results, err := geocodeAddress(client, address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, results)
	}
}

func geocodeAddress(client *maps.Client, address string) ([]maps.GeocodingResult, error) {
	req := &maps.GeocodingRequest{
		Address: address,
	}
	return client.Geocode(context.Background(), req)
}
