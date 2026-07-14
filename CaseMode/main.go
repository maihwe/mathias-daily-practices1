package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(applyCase("hello", "up"))
	fmt.Println(applyCase("WORLD", "low"))
	fmt.Println(applyCase("bridge", "cap"))
}

func applyCase(word, mode string) string {

	switch mode {
	case "up":
		return strings.ToUpper(word)

	case "low":
		return strings.ToLower(word)

	case "cap":
		return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	default:
		return word
	}

}