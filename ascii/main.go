package main

import (
	"fmt"
	"os"
	"strings"
)
func main()  {
	if len(os.Args) != 3 {
		fmt.Println("invale input: go run . string bannerfile")
		return
	}

	input := os.Args[1]
	banners := os.Args[2]

	file, err := os.ReadFile(banners)
	if err != nil {
		fmt.Println("invalid")
		return
	}

	lines := string(file)
	text := strings.Split(lines, "\n")

	for rows := 0; rows < 8; rows++ {
		for _, char := range input {
			index := int(char - 32)
			start := index * 9
			fmt.Print(text[start+rows])
		}
		fmt.Println()
	}
}