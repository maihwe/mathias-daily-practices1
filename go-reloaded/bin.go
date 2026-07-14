package main

import (
	"fmt"
	"strconv"
	"strings"
)

func binToDecimal(binStr string) (int64, error) {
	return strconv.ParseInt(binStr, 2, 64)
}
func cap(word string) string {

	words := strings.ToUpper(string(word[0])) + strings.ToLower(word[1:]) 
	return words
}
// Write a function that converts the last N words to uppercase.
// Example:
// ["How","are","you"], n=2 -> "How ARE YOU"
func uppercaseN(words []string, n int) string {
	if len(words) == 0 {
		return ""
	}
	start := len(words) - n
	for i := start; i < len(words); i++ {
	words[i] =	strings.ToUpper(words[i]) //
	}
	return strings.Join(words, " ")
}
// Question 5 – Detect Punctuation
// Write a function that returns true if a string is punctuation.
// Examples:
// "," -> true
// "." -> true
// "x" -> false
func punc(s string) bool {
	return strings.ContainsAny(s, ",.:;?!")
}

// Question 6 – Fix Article Spacing
// Fix spacing around punctuation marks.
// Example:
// ["hello", ",", "world", "!"] -> "hello, world!"
func fixArticles(words []string) string {
	res := strings.Join(words, " ")

	for _, p := range []string {",", ".", "?", "!", ":", ";"} {
		res = strings.ReplaceAll(res, p, string(p[0]))
	}
	return res
}
func main() {
	fmt.Println(binToDecimal("1010"))
	fmt.Println(binToDecimal("10"))
	fmt.Println(cap("iHWE"))
	fmt.Println(cap("mATHIAS"))
	fmt.Println(uppercaseN([]string{"How", "are", "you"}, 2))
	fmt.Println(punc(","))
	fmt.Println(punc("x"))
	fmt.Println(punc("."))
	fmt.Println(fixArticles([]string{"hello", ",", "world", "!"}))
}