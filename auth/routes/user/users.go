package routes

import (
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
