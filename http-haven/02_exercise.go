package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Found", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)


}
