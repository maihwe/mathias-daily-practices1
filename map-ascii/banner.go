package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	result := make(map[rune][]string)

	for i := 0; i < 95; i++ {
		start := i * 9

		if start+8 >= len(lines) {
			return nil, fmt.Errorf("invalid banner file")
		}

		result[rune(i+32)] = lines[start : start+8]
	}

	return result, nil
}