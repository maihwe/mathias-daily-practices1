package main

import (
	"fmt"
	"net/http"
)

func myMethod(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You made a %s request", r.Method)
}