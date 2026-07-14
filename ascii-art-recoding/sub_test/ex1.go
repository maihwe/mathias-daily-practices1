package main

import (
	"fmt"
	"os"
)

func mai() {
	file, err := os.ReadFile("s.txt")

	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(string(file))

}
