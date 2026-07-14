package asciiconverter

import "strings"

func StringToArt(input string) string {
	if input == "" {
		return ""
	}

	digits := map[rune][5]string{
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

	var out []string

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		rows := make([]string, 5)

		for _, r := range line {
			art, ok := digits[r]
			if !ok {
				return ""
			}

			for i := 0; i < 5; i++ {
				rows[i] += art[i]
			}
		}

		out = append(out, rows...)
	}

	return strings.Join(out, "\n") + "\n"
}