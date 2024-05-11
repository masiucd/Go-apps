package routes

import (
	"go-apps/auth.com/db"
	"go-apps/auth.com/input"
	"go-apps/auth.com/lib"
	"go-apps/auth.com/model"
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formValue := lib.GetFormValue(r)

	firstName := formValue("firstname")
	lastName := formValue("lastname")
	email := formValue("email")
	password := formValue("password")

	if !lib.AllFieldsValid(firstName, lastName, email, password) {
		lib.ExecuteTemplateWithData("error", w, "All fields are required")
		return
	}

	sql := db.DB
	var user model.UserRecord
	result := sql.Where("email = ?", email).First(&user)
	if result.Error != nil {
		lib.ExecuteTemplateWithData("error", w, "User already exists")
		return
	}

	hashedPassword, err := lib.HashPassword(password)
	if err != nil {
		lib.ExecuteTemplateWithData("error", w, "Error creating user")
		return
	}
	input := input.UserInput{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPassword,
	}

	error := persistence.InsertUser(input)
	if error != nil {
		lib.ExecuteTemplateWithData("error", w, "Error creating user")
		return
	}
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)

}
