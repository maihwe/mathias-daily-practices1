package main

import (
	"bufio"
	"os"
)

// LoadBanner opens a banner file and splits it into a slice of lines
func LoadBanner(style string) ([]string, error) {
	filename := style + ".txt"

	file, err := os.Open(filename)
	if err != nil {
		return nil, err // If file doesn't exist, return error
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Err()
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// GenerateASCII takes the input text and the file lines, then returns the final art string
func GenerateASCII(text string, lines []string) string {
	// 1. Handle empty input gracefully
	if text == "" {
		return ""
	}

	result := ""

	// 2. Loop through each line of the user's input text
	// (This supports basic multi-line inputs if they press Enter)
	// For simplicity, let's process the string directly character by character

	// We need 8 rows of output for our ASCII characters
	for row := 0; row < 8; row++ {
		for i := 0; i < len(text); i++ {
			asciiVal := int(text[i])

			// Only process printable ASCII characters (space 32 to tilde 126)
			if asciiVal >= 32 && asciiVal <= 126 {
				// Calculate the starting line index for this specific character in the file
				// Each character takes up 9 lines in the file (1 blank line + 8 art lines)
				charStartLine := (asciiVal-32)*9 + 1

				// Grab the specific row line for this character
				targetLine := charStartLine + row

				// Append it to our current row string
				result += lines[targetLine]
			}
		}
		// Add a newline after finishing the row for all characters
		result += "\n"
	}

	return result
}
