package main 

import (
	"fmt"
	"os"
)
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: go run . <string>")
		return
	}
	input := os.Args[1]

	letters := map[rune]string {
		'A': "Apple",
		'B': "Ball",
		'C': "Cat",
	}
	 for _, char := range input {
		value, exist := letters[char]
		if !exist {
			fmt.Printf("Error: %c Not found\n", char)
		}
		fmt.Println(value)
	}
}