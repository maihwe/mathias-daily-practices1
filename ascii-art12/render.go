package main

import "fmt"

func PrintASCIIArt(inputSegement []string, charmap map[rune][]string) {
	for _, segement := range inputSegement {

		if len(segement) == 0 {
			fmt.Println()
			continue
		}
		
		for row := 0; row < 8; row++ {

			for i := 0; i < len(segement); i++ {

				char := rune(segement[i])

				charLine := charmap[char]

				fmt.Print(charLine[row])
			}
			fmt.Println()
		}
	}

}