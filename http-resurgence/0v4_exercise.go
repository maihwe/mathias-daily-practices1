package main

import (
	"fmt"
	"net/http"
)

func DetectForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	language := r.FormValue("language")

	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}

	if language == "" {
		http.Error(w, "language is required", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s you are coding in %s!", username, language)

}
