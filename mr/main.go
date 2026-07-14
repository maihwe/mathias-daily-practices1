package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run . 'your text'")
        return
    }

    input := os.Args[1]

    banner, err := LoadBanner("standard.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    chunks := strings.Split(input, "\\n")

    for _, word := range chunks {
        if word == "" {
            fmt.Println()
            continue
        }

        for i := 0; i < 8; i++ {
            for _, char := range word {
                if char < 32 || char > 126 {
                    continue
                }
                fmt.Print(banner[char][i])
            }
            fmt.Println()
        }
    }
}
