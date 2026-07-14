package ascii

import (
	"strings"
)

func StringToArt(input string) string {
	if input == "" {
		return ""
	}

	digits := map[rune][]string {
		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			" ___ ",
			"    |",
			" ___|",
			"||   ",
			"|___ ",
		},
	}
	var output []string

	for _, lines := range strings.Split(input, "\n") {
		if lines == "" {
			continue
		}

		rows := make([]string, 5)

		for _, ch := range lines {
			art, ok := digits[ch]
			if !ok {
				return ""
			}
			for i := 0; i <5; i++ {
				rows[i] = art[i]
			}
		}
		output = append(output, rows...)
	}
	return strings.Join(output, "\n") + "\n"
}