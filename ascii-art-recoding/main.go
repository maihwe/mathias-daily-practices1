package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}

	rawInput := args[0]
	bannerName := "standard"

	if len(args) == 2 {
		bannerName = args[1]
	}

	bannerName = strings.TrimSuffix(bannerName, ".txt")
	fileName := bannerName + ".txt"

	banner, err := LoadBanner(fileName)
	if err != nil {
		fmt.Printf("Error: could not find or read banner '%s'\n", fileName)
		os.Exit(1)
	}

	// Filter out unsupported characters
	var filtered strings.Builder
	for i := 0; i < len(rawInput); i++ {
		r := rune(rawInput[i])
		
		// Handle literal \n sequences first
		if r == '\\' && i+1 < len(rawInput) && rawInput[i+1] == 'n' {
			filtered.WriteString("\\n")
			i++
			continue
		}

		// Only add if it exists in our banner map
		if _, ok := banner[r]; ok {
			filtered.WriteRune(r)
		}
	}

	finalInput := filtered.String()
	if finalInput == "" {
		return
	}

	fmt.Print(GenerateArt(finalInput, banner))
}
