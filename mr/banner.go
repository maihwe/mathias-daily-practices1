package main

import (
    "os"
    "strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(data), "\n")

    banner := make(map[rune][]string)

    for i := 0; i < 95; i++ {
        char := rune(32 + i)
        start := i * 9
        chunk := lines[start : start+8]
        banner[char] = chunk
    }

    return banner, nil
}
