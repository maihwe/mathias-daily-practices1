package ascii

import (
	"strings"
	"unicode"
)

func GenerateFont() map[rune][]string {
	fontMap := make(map[rune][]string)
	for ch := rune(32); ch <= 126; ch++ {
		fontMap[ch] = generateChar(ch)
	}
	return fontMap
}

func generateChar(ch rune) []string {
	grid := make([]string, 8)

	switch {
	case ch == ' ':
		for i := range grid {
			grid[i] = "        "
		}
	case unicode.IsDigit(ch):
		for i := range grid {
			row := []rune("........")
			if i == 0 || i == 7 {
				for j := range row {
					row[j] = '*'
				}
			} else {
				row[0], row[7] = '*', '*'
			}
			grid[i] = string(row)
		}
	case unicode.IsLetter(ch):
		if isVowel(ch) {
			grid = []string {
				"  ****  ",
				" *....* ",
				"*......*",
				"*......*",
				"*......*",
				"*......*",
				" *....* ",
				"  ****  ",
			}
		} else {
			grid = []string {
				"********",
				"*......*",
				"*......*",
				"********",
				"*......*",
				"*......*",
				"*......*",
				"********",
			}
		}
	default:
		for i := range grid {
			row := []rune ("........")
			row[0], row[7-i] = '*', '*' 
			grid[i] = string(row)
		}
	}
	return grid
}

func isVowel(ch rune) bool {
	return strings.ContainsRune("aeiou", unicode.ToLower(ch))
}