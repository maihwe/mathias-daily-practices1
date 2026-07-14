package main

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2")
}
