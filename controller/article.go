package controller

import (
	"html/template"
	"net/http"
)

type article struct {
	articleTemplate *template.Template
}

func (a article) registerRoutes() {
	http.HandleFunc("/", a.handleRequest)
}

func (a article) handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	err := a.articleTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}