package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func uppercaseWord(s string) string {
	return strings.ToUpper(s)
}

func lowercaseWord(s string) string {
	return strings.ToLower(s)
}

func capitalizeFirstLetter(s string) string {
	words := strings.Fields(s)

	for i, w := range words {
		words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
	}
	return strings.Join(words, " ")
}

func titleWord(s string) string {
	words := strings.Fields(strings.ToLower(s))

	for i, w := range words {

		if i == 0 {
			words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
			continue
		}

		if len(w) == 3 {
			words[i] = strings.ToLower(w)
		}

		if len(w) > 3 {
			words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])

		}

	}

	return strings.Join(words, " ")
}

func snakeCase(s string) string {
	s = strings.ToLower(s)
	return strings.Join(strings.Fields(s), "_")
}

func reverseWord(s string) string {

	word := strings.Fields(s)

	for i, w := range word {

		h := []rune(w)

		for a, b := 0, len(w)-1; a < b; a, b = a+1, b-1 {
			h[a], h[b] = h[b], h[a]
		}
		word[i] = string(h)
	}
	return strings.Join(word, " ")
}

var history []string

func addHistory(cmd string) string {
	if len(history) == 5 {
		history = history[1:]
	}
	history = append(history, cmd)
	return strings.Join(history, "\n")
}

func showHistory() {
	if len(history) == 0 {
		fmt.Println("\033[31;1m No history yet!\033[0m")
		return
	}

	fmt.Println("\033[32;1m<-----History----->\033[0m")
	for i, h := range history {
		fmt.Printf("%d. %s\n", i+1, h)
	}
}

func countText(text string) string {
	letters := 0
	for _, r := range text {
		if unicode.IsLetter(r) {
			letters++
		}
	}
	words := len(strings.Fields(text))
	spaces := strings.Count(text, " ")
	out := fmt.Sprintf("chars:%d letters:%d words:%d spaces:%d",
		len(text), letters, words, spaces)

	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(" ")

	fmt.Println("\033[34;1mWELCOME TO SENTINEL STRING TRANSFORMER — ONLINE\033[0m")

	fmt.Println("For Smooth Xperience, here is a quick guide for you!")
	fmt.Println("SST is a string transforming and formatting tool.\n1: The following keywords represent each transformer; upper, lower, cap, title, snake, reverse, and count.\n2: Typing any of the keywords followed by the string allows you to transform that text.")

	for {

		fmt.Println(" ")
		fmt.Println("\033[32;1mEnter the command and the string.\033[0m")

		scanner.Scan()
		scanner.Err()
		ben := scanner.Text()
		ben = strings.TrimSpace(ben)

		if ben == "history" {
			showHistory()
			continue
		}

		if ben == "exit" {
			fmt.Println("\033[32;1mGoodbye, see you again...............!\033[0m")
			break
		}

		if ben == "" {
			fmt.Println("\033[31;1mError: Empty input\033[0m")
			continue
		}

		text := strings.Fields(ben)
		sgh := strings.Join(text[1:], " ")

		if text[0] != "upper" &&
			text[0] != "lower" &&
			text[0] != "cap" &&
			text[0] != "title" &&
			text[0] != "snake" &&
			text[0] != "reverse" &&
			text[0] != "count" &&
			text[0] != "history" {
			fmt.Println("\033[31;1m Error: Unrecognised input.\033[0m")
			// continue
			fmt.Println(" ")
			fmt.Print("\033[32;1m Here are the valid commands, then follows by the string.\033[0m\n--> upper\n--> lower\n--> cap\n--> title\n--> snake\n--> reverse\n--> count\n")
			continue
		}

		if len(text) < 2 {
			fmt.Println("\033[31;1m Error: Not enough input!\033[0m")
			continue
		}

		switch text[0] {
		case "upper":
			result := uppercaseWord(sgh)
			addHistory(result)
			fmt.Println("Result:", result)

		case "lower":
			result := lowercaseWord(sgh)
			addHistory(result)
			fmt.Println("Result:", result)

		case "cap":
			result := capitalizeFirstLetter(sgh)
			addHistory(result)
			fmt.Println("Result:", result)

		case "title":
			result := titleWord(sgh)
			addHistory(result)
			fmt.Println("Result:", result)

		case "snake":
			result := snakeCase(sgh)
			addHistory(result)
			fmt.Println("Result:", result)

		case "reverse":
			result := reverseWord(sgh)
			addHistory(result)
			fmt.Println("Result:", result)

		case "count":
			result := countText(sgh)
			addHistory(result)
			fmt.Println("Result:", result)
		case "history":
			showHistory()

		default:
			fmt.Println("\033[31;1m Error: Unknown command.\033[0m")
			// continue
		}
	}

}
