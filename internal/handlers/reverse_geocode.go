package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"googlemaps.github.io/maps"
)

func ReverseGeocodeHandler(client *maps.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		lat := c.Query("lat")
		lng := c.Query("lng")
		if lat == "" || lng == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and longitude are required"})
			return
		}

		results, err := reverseGeocode(client, lat, lng)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, results)
	}
}

func reverseGeocode(client *maps.Client, lat, lng string) ([]maps.GeocodingResult, error) {
	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil, err
	}
	longitude, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		return nil, err
	}
	req := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: latitude,
			Lng: longitude,
		},
	}
	return client.ReverseGeocode(context.Background(), req)
}
