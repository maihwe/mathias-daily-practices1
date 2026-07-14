package main

import (
	"fmt"
	"os"
)

type Character struct {
	word   string
	visual []string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Errot: please provide words")
		return
	}

	input := os.Args[1]

	letters := map[rune]Character{
		'A': {
			word: "Apple",
			visual: []string{
				"  A  ",
				" A A ",
				"AAAAA",
			},
		},
		'B': {
			word: "Ball",
			visual: []string{
				"BBBB ",
				"B   B",
				"BBBB ",
			},
		},
		'C': {
			word: "Cat",
			visual: []string{
				"CCCC ",
				"C    ",
				"CCCC ",
			},
		},
	}

	for _, char := range input {
		charStr, exist := letters[char]
		if !exist {
			fmt.Printf("Error: %c Not found\n", char)
			continue
		}
		fmt.Printf("word: %s\n", charStr.word)

		fmt.Println("Visual Art :")

		for _, lines := range charStr.visual {
			fmt.Println(lines)
		}
	}
	fmt.Println()
}
