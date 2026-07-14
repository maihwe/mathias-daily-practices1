package main

import "net/http"

func DashBoard(w http.ResponseWriter, r *http.Request) {
	secret := "secret123"
	Key := r.Header.Get("X-API-Key")

	if Key != secret {
		http.Error(w, "You are not authorized to enter this page", http.StatusUnauthorized)
		return
	} else {
		http.Error(w, "Welcome", http.StatusAccepted)
	}
}
