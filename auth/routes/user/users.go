package routes

import (
	"go-apps/auth.com/db"
	"net/http"
	"text/template"
)

func UserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user, err := db.User(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.New("user.html").ParseFiles("static/user.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "user.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10"
	}

	users, err := db.Users(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templ, err := template.New("users.html").ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templ.ExecuteTemplate(w, "users.html", users)

}
