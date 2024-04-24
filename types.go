package main

import "github.com/Asaadziad/pokedex/internal/pokecache"

// auto generated
type LocationAreas struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next     string
	Previous *string
	Cache    *pokecache.Cache
}
