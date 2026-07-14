package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// quest 1
	if len(os.Args) != 3 {
		fmt.Println("Error: go run . inputfile.txt outputfile.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]
	// quest 2
	data, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error: reading file")
		return
	}
	result := string(data)

	words := strings.Fields(result)

	for i := 0; i < len(words); i++ {

		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words[i] = ""

		}

		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words[i] = ""
		}

		if words[i] == "(cap)" {
			words[i-1] = strings.ToUpper(string(words[i-1][0])) + strings.ToLower(words[i-1][1:])
			words[i] = ""
		}

		if words[i] == "(hex)" {
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
				words[i] = ""
			}
		}
		if words[i] == "(bin)" {
			n, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
				words[i] = ""
			}
		}

		if (words[i] == "a" || words[i] == "A" ) && i+1 < len(words) && 
		strings.ContainsRune("aeiouAEOIU", rune(words[i+1][0])) {
			if words[i] == "a" {
				words[i] = "an"
			} else {
				words[i] = "An"
			}
		}
		if strings.HasPrefix(words[i], "(up,") {
			s := strings.TrimPrefix(words[i], "(up,")
			s = strings.TrimSuffix(s, ")")
			s = strings.TrimSpace(s)
			n, _ := strconv.Atoi(s)
			for j := i - n; j < i; j++ {
				if j >= 0 {
					words[j] = strings.ToUpper(words[j])
				}
			}
			words[i] = ""
		}

	}
	words = strings.Fields(strings.Join(words, " "))

	result = strings.Join(words, " ")

	result = strings.ReplaceAll(result, " ,", ",")
	result = strings.ReplaceAll(result, " .", ".")
	result = strings.ReplaceAll(result, " !", "!")
	result = strings.ReplaceAll(result, " ?", "?")
	result = strings.ReplaceAll(result, " ;", ";")
	result = strings.ReplaceAll(result, " :", ":")

	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error: Writing file")
		return
	}
	fmt.Printf("Successully processed: %s : %s\n", inputfile, outputfile)
}