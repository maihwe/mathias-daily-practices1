package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func statusCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code parameter is required", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(code)

	if err != nil {
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
		return
	}
	if num < 100 || num > 599 {
		http.Error(w, "code must be a valid HTTP status code (100 – 599)", http.StatusBadRequest)
		return
	}
	w.WriteHeader(num)
	fmt.Fprintf(w, "Responding with status %d", num)
}
