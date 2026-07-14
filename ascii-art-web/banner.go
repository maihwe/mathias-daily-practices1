package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(word string) (map[rune][]string, error) {
	file, err := os.ReadFile(word)

	if err != nil {
		return nil, fmt.Errorf("Error: no such file or directory")
	}

	data := string(file)

	if len(data) == 0 {
		return nil, fmt.Errorf("Error: banner file is empty")
	}
	data = strings.ReplaceAll(data, "\r\n", "\n")
	lines := strings.Split(data, "\n")

	if len(lines) < 855 {
		return nil, fmt.Errorf("Error: banner file is incomplete or corrupted")
	}

	mapping := make(map[rune][]string)

	for i := 32; i <= 126; i++ {
		start := (i - 32) * 9
		if start+9 > len(lines) {
			return nil, fmt.Errorf("banner  is more than required")
		}
		mapping[rune(i)] = lines[start+1 : start+9]
	}

	return mapping, nil
}
