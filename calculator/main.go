package main

import (
	"fmt"
	"math"
	"strconv"
)

func calculate(a float64, op string, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("error: division by zero")
		}
		return a / b, nil
	case "%":
		if b == 0 {
			return 0, fmt.Errorf("error: modulo by zero")
		}
		return math.Mod(a, b), nil
	default:
		return 0, fmt.Errorf("error: unknown operator '%s'", op)
	}
}

func main() {
	var firstToken string
	var op string
	var b float64

	fmt.Println("===== Mathias Calculator =====")
	fmt.Println("Operators: + - * / %")
	fmt.Println("=========================")

	for {
		fmt.Print("\nEnter expression (or 'q' to quit): ")

		fmt.Scan(&firstToken)

		if firstToken == "q" {
			fmt.Println("Goodbye dear!")
			break
		}

		a, err := strconv.ParseFloat(firstToken, 64)
		if err != nil {
			fmt.Println("Invalid number. Try again.")
			continue
		}

		fmt.Scan(&op)
		fmt.Scan(&b)

		result, err := calculate(a, op, b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%.4g %s %.4g = %.4g\n", a, op, b, result)
		}
	}
}
