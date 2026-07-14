package main

import (
	"fmt"
	"strings"

	// "strings"
	"os"
)

func man() {
	file, err := os.ReadFile("s.txt")

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	result := []string{}

	content := string(file)

	parts := strings.Split(content, ",")



	for i := 0; i < len(parts); i++ {
		result = append(result, parts[i])

	}
	fmt.Println(strings.Join(result, "\n"))
	fmt.Println(len(parts))

}
