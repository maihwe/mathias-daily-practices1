package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    fontFiles := map[string]string{
        "standard":   "standard.txt",
        "shadow":     "shadow.txt",
        "thinkertoy": "thinkertoy.txt",
    }

    if len(os.Args) == 2 && os.Args[1] == "--list" {
        fmt.Println("Available fonts:")
        for name := range fontFiles {
            fmt.Println("-", name)
        }
        return
    }

    // Allow 2 or 3 args
if len(os.Args) < 2 || len(os.Args) > 3 {
    fmt.Println("Usage: go run . [STRING] [FONT]")
    return
}
input := os.Args[1]
fontName := "standard" // default font
if len(os.Args) == 3 {
    fontName = os.Args[2] // override if provided
}

    path, ok := fontFiles[fontName]
    if !ok {
        fmt.Println("Error: font not available", fontName)
        return
    }

    fontMap, err := loadFont(path)
    if err != nil {
        fmt.Println("Error loading font:", err)
        return
    }

    renderText(input, fontMap)
}
func loadFont(path string) (map[rune][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\r")
		lines = append(lines, line)
	}
	fontMap := make(map[rune][]string)

	for i := 0; i < 95; i++ {
		char := rune(i + 32)
		start := i * 9
		fontMap[char] = lines[start+1 : start+9]
	}
	return fontMap, scanner.Err()
}

func renderText(text string, fontMap map[rune][]string) {
	chunk := strings.Split(text, "\\n")

	for _, word := range chunk {
		if word == "" {
			fmt.Println()
			continue
		}
		for row := 0; row < 8; row++ {
			for _, char := range word {
				if char < 32 || char > 126 {
					continue
				}
				fmt.Print(fontMap[char][row])
			}
			fmt.Println()
		}
	}

}

