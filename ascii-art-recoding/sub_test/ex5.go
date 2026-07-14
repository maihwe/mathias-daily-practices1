package main

import (
	"fmt"
	"os"
	"strings"
)

func m()  {
	k := os.Args[1]
	file, err := os.ReadFile(k)
	if err != nil {
		fmt.Println("Error", err)
	}
	content := string(file)

	p  := []string{
		"/ \"
		"/   \"

	}
	
	for row := 0; row < 8; row ++ {
		
	}
}