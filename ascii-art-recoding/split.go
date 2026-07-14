package main

import (
	"strings"
)

func SplitInput(input string) []string {
	// We use `\n` to represent the literal backslash and 'n'
	// Alternatively, you can use "\\n"
	return strings.Split(input, `\n`)
}
