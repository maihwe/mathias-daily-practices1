package main

import (
	"fmt"
	"os"
)

func render(chunks []string, banner map[rune][]string) {
	for _, word := range chunks {

		if word == "" {
   			 fmt.Println()
    			continue
			}

		for _, char := range word {
			if char < 32 || char > 126 {
				fmt.Fprintf(os.Stderr, "Warning: skipping '%c'\n", char)
			}
		}

		for row := 0; row < 8; row++ {
			for _, char := range word {
				if char < 32 || char > 126 {
					continue
				}
				fmt.Print(banner[char][row])
			}
			fmt.Println()
		}
	}
}