package lib

import (
	"fmt"
	"net/http"
	"text/template"
)

func ExecuteTemplate(name string, w http.ResponseWriter) {
	templ, err := template.New(fmt.Sprintf("%s.html", name)).ParseFiles(fmt.Sprintf("static/%s.html", name))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = templ.ExecuteTemplate(w, fmt.Sprintf("%s.html", name), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ExecuteTemplateWithData(name string, w http.ResponseWriter, data interface{}) {
	templ, err := template.New(fmt.Sprintf("%s.html", name)).ParseFiles(fmt.Sprintf("static/%s.html", name))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = templ.ExecuteTemplate(w, fmt.Sprintf("%s.html", name), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
