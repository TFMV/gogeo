package handlers

import (
	"googlemaps.github.io/maps"
)

func NewMapsClient(apiKey string) (*maps.Client, error) {
	return maps.NewClient(maps.WithAPIKey(apiKey))
}
