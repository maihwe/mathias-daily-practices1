package main

import (
	"os"
	"strings"
)

func LoadBanner(fileName string) (map[rune][]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	filecontent := string(data)
	lines := strings.Split(filecontent, "\n")

	charmap := make(map[rune][]string)
	currentchar := rune(32)

	for i := 0; i < len(lines); i += 9 {
		if i+8 > len(lines) {
			break
		}

		// 1. Clean the 8 lines first
		cleanedLines := make([]string, 8)
		maxWidth := 0
		for row := 0; row < 8; row++ {
			cleanedLines[row] = strings.TrimSuffix(lines[i+row], "\r")
			// Track the longest line width in this block
			if len(cleanedLines[row]) > maxWidth {
				maxWidth = len(cleanedLines[row])
			}
		}

		// 2. Pad any short lines with spaces so they match maxWidth perfectly
		for row := 0; row < 8; row++ {
			if len(cleanedLines[row]) < maxWidth {
				padding := maxWidth - len(cleanedLines[row])
				cleanedLines[row] = cleanedLines[row] + strings.Repeat(" ", padding)
			}
		}

		// 3. Save to map
		charmap[currentchar] = cleanedLines
		currentchar++
	}

	return charmap, nil
}
