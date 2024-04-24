package main

import (
	"strings"
)

func Parse(text string) []string {
	splitted := strings.Split(text, " ")
	return splitted
}
