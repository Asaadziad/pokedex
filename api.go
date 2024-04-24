package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetLocations() (*LocationAreas, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var locAreas LocationAreas
	err = json.Unmarshal(body, &locAreas)
	if err != nil {

		return nil, err
	}
	return &locAreas, nil
}
