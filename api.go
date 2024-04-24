package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func bytesToLocAreas(bytes []byte) (*LocationAreas, error) {
	var locAreas LocationAreas
	err := json.Unmarshal(bytes, &locAreas)
	if err != nil {
		return nil, err
	}
	return &locAreas, nil
}

func GetLocations(cfg *Config) (*LocationAreas, error) {
	entry, ok := cfg.Cache.Get(cfg.Next)
	if ok {
		return bytesToLocAreas(entry)
	}

	res, err := http.Get(cfg.Next)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	cfg.Cache.Add(cfg.Next, body)
	return bytesToLocAreas(body)
}
