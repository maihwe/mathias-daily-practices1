package main

import (
	"fmt"
	"net/http"
)

func UserAgent(w http.ResponseWriter, r *http.Request) {
	Userinfo := r.Header.Get("User-Agent")

	if Userinfo == "" {
		http.Error(w, "userinfo missing", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "You are visiting us using: %s", Userinfo)
}
