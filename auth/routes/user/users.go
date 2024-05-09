package routes

import (
	"go-apps/auth.com/persistence"
	"net/http"
	"text/template"
)

func UserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user, err := persistence.User(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10"
	}
	users, err := persistence.Users(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templ, err := template.New("users.html").ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = templ.ExecuteTemplate(w, "users.html", users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
