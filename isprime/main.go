package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {

	// words count

	words := "the cat sat on the mat and the cat"

	text := strings.Fields(words)

	for i := 0; i < len(text); i++ {
	}
	count := make(map[string]int)

	for _, ch := range text {
		count[ch]++
	}
	for char, coun := range count {
		fmt.Printf("%s -> %d\n", string(char), coun)
	}
	fmt.Println(isPrime(2))
	fmt.Println(isPrime(4))
	fmt.Println(isPrime(7))
	fmt.Println(isPrime(9))
	fmt.Println(isPrime(19))
	prime := findPrimes(20)
	printPrime(prime)
	fmt.Println()
	restult, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(restult)
	}
	// file reader
	file, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error: unable to read the file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Err()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 5 {
			fmt.Println(line)
		}
	}
	// http
	http.HandleFunc("/", notFoundHander)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("server is running on 8080...")
	http.ListenAndServe(":8080", nil)
}

func isPrime(n int) bool {
	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func findPrimes(max int) []int {
	primes := []int{}
	for i := 2; i <= max; i++ {

		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
func printPrime(primes []int) {

	for _, pr := range primes {
		fmt.Printf("%d ", pr)
	}
}

// function that return two result
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("It cannot divde by zero")
	}
	return a / b, nil
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func notFoundHander(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Not Found", http.StatusNotFound)

}
func formHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `<form action="/greet" method="POST">
    	<input type="text" name="name" placeholder="Enter your name">
    	<button type="submit">Submit</button>
	</form>`)
}
func greetHandler(w http.ResponseWriter, r *http.Request) {
	type pageData struct {
		Name string
	}
	r.ParseForm()
	name := r.FormValue("name")
	data := pageData{Name: name}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
