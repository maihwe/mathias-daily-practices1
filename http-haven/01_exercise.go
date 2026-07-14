package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Methode Not Allowed", http.StatusMethodNotAllowed)
	}
	fmt.Fprintln(w, "pong")
}
