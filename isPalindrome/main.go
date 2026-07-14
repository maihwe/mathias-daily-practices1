package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
	runes := []rune(input)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return input == string(runes)
}

func charFrequency(s string) map[rune]int {
	frequency := make(map[rune]int)

	for _, ch := range s {

		frequency[ch]++
	}
	return frequency
}

func removeDuplicate(s string) string {
	seen := make(map[rune]bool)
	result := []rune{}

	for _, ch := range s {
		if !seen[ch] {
			seen[ch] = true
			result = append(result, ch)
		}
	}
	return string(result)
}

func isValidEmail(email string) bool {
	if !strings.Contains(email, "@") {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	if len(parts[0]) == 0 {
		return false
	}

	if !strings.Contains(parts[1], ".") {
		return false
	}
	return true
}
func main() {
	fmt.Println(isPalindrome("Mathias"))
	fmt.Println(isPalindrome("racecar"))
	fmt.Println(isPalindrome("Hello"))
	fmt.Println(isPalindrome("madam"))
	freq := charFrequency("hello world")

	for ch, count := range freq {
		fmt.Printf("%c %d\n", ch, count)
	}

	fmt.Println(removeDuplicate("hello"))
	fmt.Println(removeDuplicate("Mathias"))
	fmt.Println(removeDuplicate("mississippi"))
	fmt.Println(isValidEmail("ihwemathias@gmail.com"))
	fmt.Println(isValidEmail("mathias"))
	fmt.Println(isValidEmail("@gmail.com"))
	fmt.Println(isValidEmail("mathias@"))
	fmt.Println(isValidEmail("mathias@gmail"))

}
