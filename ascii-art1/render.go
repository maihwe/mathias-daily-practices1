package main

import (
	"strings"
)



// func RenderLine(userInput string, loadBanner map[rune][]string) []string {
// 	words := strings.Split(userInput, "\n")

// 	container := []string{}

// 	for _, word := range words {
// 		for j := 0; j < 8; j++ {
// 			for _, ch := range word {
// 				container = append(container, loadBanner[ch][j])
// 			}
// 			container = append(container, "\n")
// 		}
// 	}
// 	return container
// }


func RenderLine(input string, banner map[rune][]string) []string {
	output := make([]string, 8)

	for i := range 8 {
		var builder strings.Builder

		for _, ch := range input {
			art, ok := banner[ch]

			if !ok {
				art = banner[' ']
			}

			builder.WriteString(art[i])
		}

		output[i] = builder.String()
	}

	return output
}


