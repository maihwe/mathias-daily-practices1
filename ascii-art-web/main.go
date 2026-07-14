package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type PageData struct {
    Result string
}

func parseTemplate(w http.ResponseWriter) *template.Template {
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Templates not found", http.StatusNotFound)
        return nil
    }
    return tmpl
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/ascii-art", asciiHandler)
    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl := parseTemplate(w)
    if tmpl == nil {
        return
    }
    data := PageData{Result: ""}
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, "Template execution error", http.StatusInternalServerError)
    }
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
    text := r.FormValue("text")
    banner := r.FormValue("banner")
    if banner == "" {
        banner = "standard"
    }

    bannerFile := banner + ".txt"
    bannerMap, err := LoadBanner(bannerFile)
    if err != nil {
        http.Error(w, "Banner not found", http.StatusNotFound)
        return
    }

    result := GenerateArt(text, bannerMap)
    tmpl := parseTemplate(w)
    if tmpl == nil {
        return
    }
    data := PageData{Result: result}
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, "Template execution error", http.StatusInternalServerError)
    }
}
