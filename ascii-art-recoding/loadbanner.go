package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(s string) (map[rune][]string, error) {
	file, err := os.ReadFile(s)

	if err != nil {
		return nil, fmt.Errorf("Error: failed to read file")
	}
	mapping := make(map[rune][]string)

	content := string(file)

	if len(content) == 0 {
		return nil, fmt.Errorf("empty file")
	}

	lines := strings.Split(content, "\n")
	if len(lines) < 855 {
		return nil, fmt.Errorf("bad file format")
	}

	for i := 32; i <= 126; i++ {
		start := (i - 32) * 9
		if start > len(lines) {
			break
		}
		mapping[rune(i)] = lines[start+1 : start+9]
	}
	return mapping, nil
}
