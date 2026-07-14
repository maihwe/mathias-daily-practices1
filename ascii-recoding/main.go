package main

import (
	"fmt"
	"os"
)

func main() {
	// ParseArgs validates everything in one call:
	//   1. argument count (2 or 3)
	//   2. input string   (not empty, not only \n markers)
	//   3. font name      (must be standard | shadow | thinkertoy)
	input, fontPath, err := ParseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Generate loads the banner file and prints the ASCII art.
	if err := Generate(input, fontPath); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
