package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	// Empty input
	if input == "" {
		return ""
	}

	// Convert literal \n into real newline
	input = strings.ReplaceAll(input, `\n`, "\n")

	// Special case: input contains only newlines
	onlyNewlines := true
	for _, ch := range input {
		if ch != '\n' {
			onlyNewlines = false
			break
		}
	}
	if onlyNewlines {
		return input
	}

	parts := strings.Split(input, "\n")
	var result strings.Builder

	for _, part := range parts {
		// Blank line between words
		if part == "" {
			result.WriteString("\n")
			continue
		}

		// Print 8 rows of banner text
		for row := 0; row < 8; row++ {
			for _, ch := range part {
				result.WriteString(banner[ch][row])
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}