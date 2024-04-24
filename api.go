package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetLocations(cfg *Config) (*LocationAreas, error) {
	res, err := http.Get(cfg.Next)
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
