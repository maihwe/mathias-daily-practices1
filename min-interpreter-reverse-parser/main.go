package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var stack []int

	fmt.Println("Enter RPN expression (e.g., '10 5 +') or 'exit' to quit:")

	for {
		if !scanner.Scan() {
			scanner.Err()
			break
		}
		te := scanner.Text()
		te = strings.TrimSpace(te)

		if te == "exit" {
			fmt.Println("Goodbye........!!")
			break
		}

		if te == "" {
			fmt.Println("Error: Empty input")
			continue
		}

		ret := strings.Fields(te)
		isValidExpression := true

		for _, r := range ret {
			if r == "-" || r == "+" || r == "*" || r == "/" {
				if len(stack) < 2 {
					fmt.Println("Error: Not enough operands in stack.")
					isValidExpression = false
					break
				}

				right := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				var result int
				switch r {
				case "+":
					result = left + right
				case "-":
					result = left - right
				case "*":
					result = left * right
				case "/":
					if right == 0 {
						fmt.Println("Error: Division by zero")
						isValidExpression = false
						break
					}
					result = left / right
				}
				
				if !isValidExpression {
					break
				}
				stack = append(stack, result)

			} else {
				// Parse Base-16 (Hexadecimal) or Base-10 (Decimal)
				num, err := parseNumber(r)
				if err != nil {
					fmt.Println("Error:", err)
					isValidExpression = false
					break
				}
				stack = append(stack, num)
			}
		}

		// Print the final result if the calculation was successful
		if isValidExpression && len(stack) > 0 {
			fmt.Printf("Result: %d\n", stack[len(stack)-1])
			// Clear stack for the next calculation line
			stack = nil 
		} else {
			// Clear stack on failure to prevent memory corruption
			stack = nil 
		}
	}
}

// Fixed helper function integrating your Base-16 / Base-10 string parsing logic
func parseNumber(s string) (int, error) {
	num := 0
	isHex := false

	// Check if token looks like a Hexadecimal number (e.g., contains A-F or prefix)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f') {
			isHex = true
			break
		}
	}

	if isHex {
		for i := 0; i < len(s); i++ {
			c := s[i]
			value := 0
			if c >= '0' && c <= '9' {
				value = int(c - '0')
			} else if c >= 'A' && c <= 'F' {
				value = int(c-'A') + 10
			} else if c >= 'a' && c <= 'f' {
				value = int(c-'a') + 10
			} else {
				return 0, fmt.Errorf("invalid hex character '%c'", c)
			}
			num = num*16 + value
		}
	} else {
		// Base-10 Decimal Parsing
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c < '0' || c > '9' {
				return 0, fmt.Errorf("invalid decimal character '%c'", c)
			}
			num = num*10 + int(c-'0')
		}
	}
	return num, nil
}
