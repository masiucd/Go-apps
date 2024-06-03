package routes

import (
	"errors"
	"fmt"
	"go-apps/auth.com/db"
	"go-apps/auth.com/input"
	"go-apps/auth.com/lib"
	"go-apps/auth.com/model"
	usersdao "go-apps/auth.com/persistence/users-dao"

	"net/http"
	"strconv"
)

func UserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user := usersdao.User(id)
	if user == nil {
		data := lib.ErrorData{
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
	users, err := usersdao.Users(limit)

	if err != nil {
		data := lib.ErrorData{
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
	// TODO - If logged in redirect to profile
	err := r.ParseForm()
	if err != nil {
		lib.ExecuteTemplateWithData("signup", w, "Error parsing form")
		return
	}

	firstname, lastname, email, password, err := validateFormValues(r)
	if err != nil {
		lib.ExecuteTemplateWithData("signup", w, err.Error())
		return
	}

	user := usersdao.UserByEmail(email)
	// if user exists, return an error
	if user != nil {
		lib.ExecuteTemplateWithData("signup", w, "User already exists")
		return
	}

	hashedPassword, err := lib.HashPassword(password)

	if err != nil {
		data := lib.ErrorData{
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

	error := usersdao.InsertUser(input)
	if error != nil {
		data := lib.ErrorData{
			Message: "Error creating user",
			Code:    http.StatusInternalServerError,
		}
		lib.ExecuteTemplateWithData("error", w, data)
		return
	}
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}

// TODO - Implement login
func Profile(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session_token")
	if err != nil {
		fmt.Println("Error getting cookie", err.Error())
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Println("c.Value", c.Value)
	// TODO move to persistence layer
	var sessionRecord model.SessionRecord
	sql := db.DB
	result := sql.Where("token = ?", c.Value).First(&sessionRecord)
	if result.RowsAffected == 0 {
		fmt.Println("Session not found in DB")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Println("Session token", c.Value, result.RowsAffected, sessionRecord.Token, sessionRecord.ID, sessionRecord.Email)
	// user := persistence.UserByEmail(sessionRecord.Email)

	// w.Write([]byte("Welcome " + user.FirstName + " " + user.LastName))
	w.Write([]byte("Welcome "))

}

func validateFormValues(r *http.Request) (string, string, string, string, error) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if isEmpty(firstname, lastname, email, password) {
		return "", "", "", "", errors.New("all fields are required")
	}

	return firstname, lastname, email, password, nil
}

func isEmpty(values ...string) bool {
	for _, value := range values {
		if value == "" {
			return true
		}
	}
	return false
}
