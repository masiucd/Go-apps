package routes

import (
	"go-apps/auth.com/input"
	"go-apps/auth.com/lib"
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
	lib.ExecuteTemplateWithData("users", w, users)
}

// Renders the page where a user can Sign up
func SignUp(w http.ResponseWriter, r *http.Request) {
	lib.ExecuteTemplate("signup", w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	lib.ExecuteTemplate("login", w)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	formValue := getFormValue(r)
	input := input.UserInput{
		FirstName: formValue("firstname"),
		LastName:  formValue("lastname"),
		Email:     formValue("email"),
		Password:  formValue("password"),
	}

	error := persistence.InsertUser(input)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)

}

func getFormValue(r *http.Request) func(key string) string {
	return func(key string) string {
		if val, ok := r.Form[key]; ok {
			return val[0]
		}
		return ""
	}
}
