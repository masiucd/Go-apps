package routes

import (
	"fmt"
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		lib.ExecuteTemplateWithData("signup", w, "Error parsing form")
		return
	}
	fmt.Println("r.Form", r.Form)

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println(
		"firstname", firstname,
		"lastname", lastname,
		"email", email,
		"password", password,
	)

	if firstname == "" || lastname == "" || email == "" || password == "" {
		lib.ExecuteTemplateWithData("signup", w, "All fields are required")
		return
	}

	exists := persistence.DoesUserExist(email)
	if exists {
		fmt.Println("DoesUserExist ", exists)
		lib.ExecuteTemplateWithData("signup", w, "User already exists")
		return
	}

	hashedPassword, err := lib.HashPassword(password)
	fmt.Println("hashedPassword", hashedPassword)
	if err != nil {
		lib.ExecuteTemplateWithData("error", w, "Error creating user")
		return
	}
	input := input.UserInput{
		FirstName: firstname,
		LastName:  lastname,
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
