package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Calculate(w http.ResponseWriter, r *http.Request) {
	A := r.URL.Query().Get("a")
	B := r.URL.Query().Get("b")
	Opt := r.URL.Query().Get("op")

	fnum, err := strconv.Atoi(A)

	if err != nil {
		http.Error(w, "string conversion failed", 400)
		return
	}

	snum, err := strconv.Atoi(B)

	if err != nil {
		http.Error(w, "string conversion failed", 500)
		return
	}

	var result int
	switch Opt {
	case "add":
		result = fnum + snum
	case "subtract":
		result = fnum - snum
	case "multiply":
		result = fnum * snum

	default:
		http.Error(w, "Unknown Operation", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Result: %d", result)
}
