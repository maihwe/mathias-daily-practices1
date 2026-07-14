package main

import (
	"fmt"
	"io"
	"net/http"
)

func Count(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case http.MethodGet:
		fmt.Fprint(w, "Send a POST request with text to count words")


	case http.MethodPost:

		defer req.Body.Close()

		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Failed to read the body", http.StatusInternalServerError)
			return
		}

		charcoun := len([]rune(string(body)))
		fmt.Fprintf(w, "%d", charcoun)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}