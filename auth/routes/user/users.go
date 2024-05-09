package routes

import (
	"fmt"
	"go-apps/auth.com/persistence"
	"net/http"
	"strconv"
	"text/template"
)

func UserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user := persistence.User(id)
	tmpl, err := template.New("user.html").ParseFiles("static/user.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.ExecuteTemplate(w, "user.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if limit == 0 {
		limit = 10
	}
	users, err := persistence.Users(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	executeTemplateWithData("users", w, users)
}

// Renders the page where a user can Sign up
func SignUp(w http.ResponseWriter, r *http.Request) {
	// get the formData from the user
	executeTemplate("signup", w)
}

func executeTemplate(name string, w http.ResponseWriter) {
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

func executeTemplateWithData(name string, w http.ResponseWriter, data interface{}) {
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
