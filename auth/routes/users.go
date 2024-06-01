package routes

import (
	"fmt"
	"go-apps/auth.com/input"
	"go-apps/auth.com/lib"
	"go-apps/auth.com/model"
	"go-apps/auth.com/persistence"
	"net/http"
	"strconv"
)

type ErrorData struct {
	Message string
	Code    int
}

func UserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user := persistence.User(id)
	if user == nil {
		data := ErrorData{
			Message: "User not found",
			Code:    http.StatusNotFound,
		}
		lib.ExecuteTemplateWithData("error", w, data)
		return
	}
	lib.ExecuteTemplateWithData("user", w, user)

}

func Users(w http.ResponseWriter, r *http.Request) {
	limitQuery := r.URL.Query().Get("limit")
	if limitQuery == "" {
		limitQuery = "10"
	}
	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	users, err := persistence.Users(limit)

	if err != nil {
		data := ErrorData{
			Message: "Error fetching users",
			Code:    http.StatusInternalServerError,
		}
		lib.ExecuteTemplateWithData("error", w, data)
		return
	}
	fmt.Println(users[0].LastName)
	data := struct {
		Users []*model.UserRecord
		Title string
	}{
		Users: users,
		Title: "Users",
	}

	lib.ExecuteTemplateWithData("users", w, data)

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

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")

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

	if err != nil {
		data := ErrorData{
			Message: "Error creating user",
			Code:    http.StatusInternalServerError,
		}
		lib.ExecuteTemplateWithData("error", w, data)
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
		data := ErrorData{
			Message: "Error creating user",
			Code:    http.StatusInternalServerError,
		}
		lib.ExecuteTemplateWithData("error", w, data)
		return
	}
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}
