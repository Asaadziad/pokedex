package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cfg := Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: nil,
	}
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	fmt.Print("pokedex > ")
	for scanner.Scan() {
		command := Parse(scanner.Text())
		err := Eval(&cfg, command)
		if err != nil {
			break
		}
		fmt.Print("pokedex > ")
	}
}
