package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Check command line arguments
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}

	banner := "standard"
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}

	fileName := banner + ".txt"

	// 2. Call our new function from banner.go
	charmap, err := LoadBanner(fileName)
	if err != nil {
		fmt.Println("Error: reading file")
		return
	}

	// 3. Prepare the input string
	inputfile := os.Args[1]
	inputSegement := strings.Split(inputfile, "\\n")

	// 4. Call our new function from render.go
	PrintASCIIArt(inputSegement, charmap)
}
