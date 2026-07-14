package main

import (
	"strings"
)
func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	input = strings.ReplaceAll(input, `\n`, "\n")

	onlyNewLine := true

	for _, mat := range input {
		if mat != '\n' {
			onlyNewLine = false
			break
		}
	}
	if onlyNewLine {
		return input
	}

	parts := strings.Split(input, "\n")
	var want strings.Builder

	for _, part := range parts {
		if part == "" {
			want.WriteString("\n")
			continue
		}

		maxrows := 0
		for _, ch := range part {
			if rows := len(banner[ch]); rows > maxrows {
				maxrows = rows
			}
		}

		for row := 0; row < maxrows; row ++ {
			for _, mat := range part {
				want.WriteString(banner[mat][row])
			}
			want.WriteString("\n")
		}
	}
	return want.String()
}