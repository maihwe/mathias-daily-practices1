package main

import (
	"fmt"
	"net/http"
)


func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

	
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	
	if name == "" {
		fmt.Fprintf(w, "Greetings, Stranger!")
	}
	fmt.Fprintf(w, "Greetings, %s!", name)
}