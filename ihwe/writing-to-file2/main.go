package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("Hello.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineNumber := 1
    for scanner.Scan() {
        fmt.Printf("Line %d: %s\n", lineNumber, scanner.Text())
        lineNumber++
    }
}
