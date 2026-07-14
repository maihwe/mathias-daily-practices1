package main

import (
	"fmt"
	"strings"
)

func main() {
	input, ok := getInput()
	if !ok {
		return
	}

	fontFile := getFontFile()

	banner, err := LoadBanner(fontFile)
	if err != nil {
		fmt.Println("Error: reading font file:", fontFile)
		return
	}

	chunks := strings.Split(input, "\\n")
	render(chunks, banner)
}