package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main()  {
	scanner := bufio.NewScanner(os.Stdin)


		fmt.Println(" ")
		fmt.Println("Welcome To My Base Converter!")
		fmt.Println("You can convert between bases such as base 2, 16 and 10.")
		fmt.Print("For Smooth Xperience, here is a quick guide.\n1: You type directly in your terminal.\n2: Your firt word must be 'convert'follows by base and str.\n3: Also, you can type quit to exit the program at any point.\n")
		fmt.Println(" ")

	for {

		scanner.Scan()
		texts := scanner.Text()
		texts = strings.TrimSpace(texts)
	

		if texts == "Quit" {
			break
		}
		if texts == "" {
			fmt.Println("Error: Empty input")
			continue
		}
	
		splitText := strings.Fields(texts)

		if len(splitText) != 3 {
			fmt.Println("Error: Not enough input!")
			continue
		}

		if splitText[0] != "convert" {
			fmt.Println("Error: The words must start with convert!")
		}

		secondPart := splitText[1]
		thirdpart := splitText[2]

		switch thirdpart {
			
		case "hex":
			n, err := strconv.ParseInt(secondPart, 16, 64)
			if err != nil {
				fmt.Println("Error:", secondPart + " is not a valid hex! ")
				continue
			}
			fmt.Println("Decimal:", n)
		case "bin":
			n, err := strconv.ParseInt(secondPart, 2, 64)
			if err != nil {
				fmt.Println("Error:", secondPart + " is not a valid binary! ")
				continue
			}
			fmt.Println("Decimal:", n)

		case "dec":
			n, err := strconv.ParseInt(secondPart, 10, 64)
			if err != nil {
				fmt.Println("Error:", secondPart + " is not a valid decimal! ")
				continue	
			}
			fmt.Printf("Binary %b\n", n)
			fmt.Printf("Hex %X\n", n)
		}

	}
}
