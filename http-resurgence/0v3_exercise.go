package main

import (
	"fmt"
	"net/http"
	
)

 func myHeader(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Custom-Token")
	if token == "" {
		fmt.Fprintf(w, "X-Custom-Token header is missing")
		return
	}
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "Content-Type Not Provided"
	}
	fmt.Fprintf(w, "Token received: %s\n", token)
	fmt.Fprintf(w, "Content-Type: %s\n", contentType)
}