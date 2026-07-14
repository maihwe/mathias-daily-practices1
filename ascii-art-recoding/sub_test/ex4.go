package main

import (
	"fmt"
	"os"
)

func main() {
	r := "ex4.txt"
	file, err := os.ReadFile(r)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	content := string(file)

	for _, ch := range content {
		index := int(ch - 32)
		t := string(ch)
		fmt.Println(t, "ascii =", ch, "fontIndex =", index)
	}
}
