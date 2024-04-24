package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	fmt.Print("pokedex > ")
	for scanner.Scan() {
		command := Parse(scanner.Text())
		err := Eval(command)
		if err != nil {
			break
		}
		fmt.Print("pokedex > ")
	}
}
