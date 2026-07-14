package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	//Step 1 -> Create and Write file
	fmt.Println("===WRITING===")
	file, err := os.Create("Hello.txt")
	if err != nil {
		fmt.Println("error", err)
		return
	}

	file.WriteString("Hello: Ihwe Mathias!\n")
	file.WriteString("You are welcome to file Handling!\n")
	file.Close()
	fmt.Println("file writen")

	//Step 2 Append
	fmt.Println("\n===APPEND===")
	file, err = os.OpenFile("Hello.txt",
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	file.WriteString("This Line was added later!\n")
	file.Close()
	fmt.Println("Line Appened")
 
	//Step  3 -> Read file
	fmt.Println("\n===READING===")
	file, err = os.Open("Hello.txt")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Err()
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}