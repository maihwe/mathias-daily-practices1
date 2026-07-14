package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/v2", Welcome)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/count", Count)
	http.HandleFunc("/calculate", Calculate)
	http.HandleFunc("/agent", UserAgent)
	http.HandleFunc("/dashboard", DashBoard)
	http.HandleFunc("/legacy", Redirect)

	fmt.Println("server is running on http://local:8080...")
	http.ListenAndServe(":8080", nil)

}

