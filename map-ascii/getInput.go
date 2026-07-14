package main

import (
	"fmt"
	"os"
	"strings"
)

func getInput() (string, bool) {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . 'text' [fontfile]")
		return "", false
	}

	input := os.Args[1]
	if input == "" {
		return "", false
	}

	stripped := strings.ReplaceAll(input, "\\n", "")
	if stripped == "" {
		fmt.Println("Error: only newlines")
		return "", false
	}

	return input, true
}

func getFontFile() string {
	if len(os.Args) == 3 {
		return os.Args[2]
	}
	return "standard.txt"
}