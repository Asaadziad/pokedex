package main

import (
	"fmt"
	"os"
)

type cmd struct {
	name        string
	description string
	callback    func() error
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
			callback:    nil,
		},
	}
}

func cmdHelp() error {
	fmt.Println("Welcome to pokedex")
	fmt.Println("Usage:")
	for _, i := range Commands() {
		fmt.Println(i.name + ": " + i.description)
	}
	return nil
}

func cmdExit() error {
	os.Exit(0)
	return nil
}

func cmdMap() error {
	return nil
}

func GetCommand(commandName string) cmd {
	return Commands()[commandName]
}
