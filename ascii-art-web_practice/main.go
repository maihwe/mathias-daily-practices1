package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Printf("Error: connot read index.html %v\n", err)
		return
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	fmt.Println("server starting on http://localhost:8081...")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("server fail to start %V\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page Not Found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Check if the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request: Method Not Allowed", http.StatusBadRequest)
		return
	}

	// 2. Extract user input values
	userText := r.FormValue("user-text")
	bannerStyle := r.FormValue("banner-style")

	// 3. Validate the input values
	if bannerStyle != "standard" && bannerStyle != "shadow" && bannerStyle != "thinkertoy" {
		http.Error(w, "400 Bad Request: Invalid banner style", http.StatusBadRequest)
		return
	}

	// ==========================================
	// THIS IS THE NEW PARTS ADDED BELOW:
	// ==========================================

	// 4. Load the file lines using ascii.go
	bannerLines, err := LoadBanner(bannerStyle)
	if err != nil {
		http.Error(w, "404 Not Found: Banner file missing", http.StatusNotFound)
		return
	}

	// 5. Calculate the ASCII art text using ascii.go
	asciiResult := GenerateASCII(userText, bannerLines)

	// 6. Return a successful response and display the art
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, asciiResult)
}
