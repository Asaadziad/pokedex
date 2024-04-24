package main

import (
	"fmt"
	"os"
)

type cmd struct {
	name        string
	description string
	callback    func(*Config) error
}

func Commands() map[string]cmd {
	return map[string]cmd{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    cmdHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit pokedex",
			callback:    cmdExit,
		},
		"map": {
			name:        "map",
			description: "gets all the location areas",
			callback:    cmdMap,
		},
		"mapb": {
			name:        "mapb",
			description: "gets the previous 20 location areas",
			callback:    cmdMapb,
		},
	}
}

func cmdHelp(cfg *Config) error {
	fmt.Println("Welcome to pokedex")
	fmt.Println("Usage:")
	for _, i := range Commands() {
		fmt.Println(i.name + ": " + i.description)
	}
	return nil
}

func cmdExit(cfg *Config) error {
	os.Exit(0)
	return nil
}

func cmdMap(cfg *Config) error {
	areas, err := GetLocations(cfg)
	if err != nil {
		return err
	}
	cfg.Next = areas.Next
	cfg.Previous = areas.Previous
	for _, e := range areas.Results {
		fmt.Println(e.Name)
	}
	return nil
}

func cmdMapb(cfg *Config) error {
	if cfg.Previous == nil {
		cfg.Next = "https://pokeapi.co/api/v2/location-area/"
	} else {
		cfg.Next = *cfg.Previous
	}
	areas, err := GetLocations(cfg)
	if err != nil {
		return err
	}
	cfg.Next = areas.Next
	cfg.Previous = areas.Previous
	for _, e := range areas.Results {
		fmt.Println(e.Name)
	}
	return nil
}

func GetCommand(commandName string) cmd {
	return Commands()[commandName]
}
