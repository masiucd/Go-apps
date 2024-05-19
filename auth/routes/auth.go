package routes

import (
	"go-apps/auth.com/lib"
	"net/http"
)

// Login renders the login page
func Login(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// formValue := lib.GetFormValue(r)
	// email := formValue("email")
	// password := formValue("password")
	// if !lib.AllFieldsValid(email, password) {
	// 	lib.ExecuteTemplateWithData("error", w, "All fields are required")
	// 	return
	// }

	// var user model.UserRecord
	// lib.ExecuteTemplateWithData("profile", w, user)
	lib.ExecuteTemplate("login", w)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formValue := lib.GetFormValue(r)
	email := formValue("email")
	password := formValue("password")
	if !lib.AllFieldsValid(email, password) {
		lib.ExecuteTemplateWithData("error", w, "All fields are required")
		return
	}

	// var user model.UserRecord
	// lib.ExecuteTemplateWithData("profile", w, user)
	lib.ExecuteTemplate("login", w)
}

// Logout logs out a user
func Logout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
