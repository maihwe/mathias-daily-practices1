package main

import (
	"fmt"
	"net/http"
)
func main() {
    apiMux := http.NewServeMux()
    apiMux.HandleFunc("/v1/ping", pingHandler)
    apiMux.HandleFunc("/v1/greet", greetHandler)

    mainMux := http.NewServeMux()
    mainMux.HandleFunc("/method-inspector", myMethod)
    mainMux.HandleFunc("/echo", echoHandler)
    mainMux.HandleFunc("/headers", myHeader)
    mainMux.HandleFunc("/form", DetectForm)
    mainMux.HandleFunc("/status", statusCode)
	mainMux.HandleFunc("/render", renderHandler)
    mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

    fmt.Println("server is running on :8080...")
    http.ListenAndServe(":8080", mainMux)
}