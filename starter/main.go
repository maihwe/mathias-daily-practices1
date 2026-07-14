package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Error: please enter a word")
		return
	}

	var data string

	if strings.HasSuffix(args[0], ".txt") {

		file, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println(err)
			return
		} 
		data = string(file)
	} else {
		data = args[0]
	}
	counts := make(map[rune]int)

	data = strings.ToLower(string(data))

	for _, r := range data {
		counts[r]++
	}

	for ch, count := range counts {
		fmt.Printf("%c -> %d\n", ch, count)
	}
}