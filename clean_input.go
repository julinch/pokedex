package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var slice []string
	text = strings.ToLower(text)
	slice = strings.Fields(text)
	return slice
}
