package main

import (
	"fmt"
	"strings"

)

var cont int
var endwith int


func ends()  {
	fmt.Println("Do you want to exit?\n1: = Yes\n2: = No")
	fmt.Scanln(&endwith)

	if endwith == 1 {
		fmt.Println("Goodbye, see you again.......!")
		return
	} else {
		main()
	}

}
 func asIfToContinue() {
	fmt.Println("Do you want to continue with another operation?:\n1: = Yes\n2: = No")
	fmt.Scanln(&cont)

		if cont == 2 {
			fmt.Println("Goodbye, see you again.......!")
			return
		} else {
			main()
		}
 }



var showHis int

var history []string

func addHistory(cmd string) {
	if len(history) == 5 {
		history = history[1:]
	}
	history = append(history, cmd)
}

func showHistory() {
	var showHis int

	fmt.Println("How many recent commands do you want to see? (1-5):")
	fmt.Scanln(&showHis)

	if showHis > len(history) {
		showHis = len(history)
	}

	

	fmt.Println("\033[1;32m---- History ----\033[0m")
	fmt.Println(strings.Join(history[len(history)-showHis:], "\n"))
	fmt.Println(strings.Join(history[showHis:], "\n"))
	fmt.Println(strings.Join(history[:showHis], "\n"))
}

func main()  {
var input1 float64
var input2 float64
var operator string
var help string


for {
	fmt.Println(" ")
	fmt.Println("Welcome To My Calculator!")
	fmt.Println("The available operations to perform are listed below!:\n1: = Divsion\n2: = Multiplication\n3: = Addition\n4: = And subtraction ")
	fmt.Println(" ")
	fmt.Println("Select an operator: ")
	fmt.Println("div, mult, add, sub, quit")
	fmt.Scanln(&operator)
	operator = strings.ToLower(strings.TrimSpace(operator))
	

	

	if operator != "div" &&
		operator != "mult" &&
		operator != "add"  &&
		operator != "sub" &&
		operator != "quit" &&
		operator != "history"{
		

		fmt.Println("Error: invalid input! Please choose from the availaible operations provided above.")
		fmt.Println(" ")
		fmt.Println("Need help? type help for guide.")
		fmt.Scanln(&help)
		fmt.Println("Please, follow the steps below:\n1: = Type the operation you want to perform from the ones provided\n2: = Enter the numbers, one after each prompt\n3: = The numbers must be non-zero digit\n4: And also letters are not accepted")
		fmt.Println(" ")
		asIfToContinue()
		break
	}


	if operator == "quit" {
		ends()
		break
	}

	if operator == "history" {
		showHistory()
		continue

	}
	addHistory(operator)
	
	fmt.Println("Enter a number")
	fmt.Scanln(&input1)
	if input1 == 0 {
		fmt.Println("Error: the number must be greater than zero and the second number must be supplied.")
		asIfToContinue()
		break
	}

	fmt.Println("Enter the second number")
	fmt.Scanln(&input2)

	if input2 == 0 {
		fmt.Println("Error: the number must be greater than zero")
		fmt.Println(" ")
		asIfToContinue()
		break
	}
	
	if operator == "div" {
		result := input1 / input2
		fmt.Println("Result:", input1 / input2)
		addHistory(fmt.Sprintf("Result: %.0f / %.0f = %.0f",  input1,  input2, result))
		fmt.Println(" ")
		asIfToContinue()
		break
	}
	if operator == "mult" {
		result := input1 * input2
		fmt.Println("Result:", input1 * input2)
		addHistory(fmt.Sprintf("Result: %.0f * %.0f = %.0f",  input1, input2, result))
		fmt.Println(" ")
		asIfToContinue()
		break
	}
	if operator == "add" {
		result := input1 + input2
		addHistory(fmt.Sprintf("Result: %.0f + %.0f = %.0f",  input1,  input2, result))

		fmt.Println(" ")
		asIfToContinue()
		break
	}
	if operator == "sub" {
		result := input1 - input2
		addHistory(fmt.Sprintf("Result: %.0f - %.0f = %.0f",  input1, input2, result))
		fmt.Println(" ")
		asIfToContinue()
		break
	}
	
}

}
	