package main

import (
	"html/template"
	"net/http"
)

const tmplStr = `
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
  <h1>{{.Title}}</h1>
  <p>{{.Body}}</p>
</body>
</html>
`

type PageData struct {
	Title string
	Body  string
}

var tmpl = template.Must(template.New("page").Parse(tmplStr))

func renderHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")

	if title == "" || body == "" {
		http.Error(w, "title and body are required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	err := tmpl.Execute(w, PageData{Title: title, Body: body})
	if err != nil {
		http.Error(w, "template execution failed", http.StatusInternalServerError)
	}
}
