package main

import (
	"fmt"
	"os"
)

func ain() {
	r := "ex3.txt"
	file, err := os.ReadFile(r)
	if err != nil {
		fmt.Println("Error", err)
	}
	content := string(file)

	for _, ch := range content {
		fmt.Println(string(ch),  ch)
	}
}
