package routes

import (
	"go-apps/auth.com/lib"
	"go-apps/auth.com/persistence"
	"net/http"
)

// Login renders the login page - GET request
func Login(w http.ResponseWriter, r *http.Request) {
	lib.ExecuteTemplate("login", w)
}

// LoginUser logs in a user - POST request
func LoginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formValue := lib.GetFormValue(r)
	email := formValue("email")
	password := formValue("password")
	if !lib.AllFieldsValid(email, password) {
		lib.ExecuteTemplateWithData("error", w, "All fields are required")
		return
	}

	user := persistence.UserByEmail(email)
	if user == nil {
		data := lib.ErrorData{
			Message: "User does not exist",
			Code:    http.StatusNotFound,
		}
		lib.ExecuteTemplateWithData("error", w, data)
		return
	}

	ok := lib.VerifyPassword(password, user.Password)
	if !ok {
		data := lib.ErrorData{
			Message: "Invalid credentials",
			Code:    http.StatusUnauthorized,
		}
		lib.ExecuteTemplateWithData("error", w, data)
		return
	}

	// TODO - Implement session management

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

// Logout logs out a user - POST request
// TODO - Implement this
func Logout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
