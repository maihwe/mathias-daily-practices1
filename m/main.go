package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, ok := getInput()
	if !ok {
		return
	}

	fontFile := getFontFile()

	lines, err := readLines(fontFile)
	if err != nil {
		fmt.Println("Error: reading font file:", fontFile)
		return
	}

	chunks := strings.Split(input, "\\n")
	render(chunks, lines)
}

func getInput() (string, bool) {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run main.go 'your text' [fontfile]")
		return "", false
	}

	input := os.Args[1]
	if input == "" {
		return "", false
	}

	stripped := strings.ReplaceAll(input, "\\n", "")
	if stripped == "" {
		fmt.Println("Error: input contains only newlines")
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

func render(chunks []string, lines []string) {
	for _, word := range chunks {
		if word == "" {
			fmt.Println()
			continue
		}

		for _, char := range word {
			if char < 32 || char > 126 {
				fmt.Fprintf(os.Stderr, "Warning: skipping unsupported character '%c'\n", char)
			}
		}

		for i := 1; i <= 8; i++ {
			for _, char := range word {
				if char < 32 || char > 126 {
					continue
				}
				startLine := int(char-32)*9 + i
				if startLine < len(lines) {
					fmt.Print(lines[startLine])
				}
			}
			fmt.Println()
		}
	}
}

func readLines(path string) ([]string, error) {
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
	return lines, scanner.Err()
}