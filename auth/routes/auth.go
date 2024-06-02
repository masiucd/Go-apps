package routes

import (
	"fmt"
	"go-apps/auth.com/db"
	"go-apps/auth.com/lib"
	"go-apps/auth.com/model"
	"go-apps/auth.com/persistence"
	"net/http"
	"time"

	"github.com/google/uuid"
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

	ok := lib.VerifyPassword(user.Password, password)
	if !ok {
		data := lib.ErrorData{
			Message: "Invalid credentials",
			Code:    http.StatusUnauthorized,
		}
		lib.ExecuteTemplateWithData("error", w, data)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(time.Minute * 60) // 1 hour

	sql := db.DB
	var session model.SessionRecord
	// TODO Before creating a new session, delete the old session if it exists
	// result := db.DB.Where("id = ?", user.ID).Delete(&model.SessionRecord{})
	result := sql.First(&session, user.ID)
	// if there a session stored with the id then we delete it before creating a new one
	fmt.Println("Session found", session, result.RowsAffected)
	if result.RowsAffected > 0 {
		fmt.Println("Deleting session", result.RowsAffected)
		sql.Delete(&session)
	}

	if result.Error != nil {
		fmt.Println("Error deleting session", result.Error)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	// store session in DB - TODO move to persistence layer
	db.DB.Create(&model.SessionRecord{
		UserID:    user.ID,
		Email:     user.Email,
		Token:     sessionToken,
		ExpiresAt: expiresAt.Unix(),
	})

	// store session in cookie
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "session_token",
	// 	Value:   sessionToken,
	// 	Expires: expiresAt,
	// })
	// fmt.Println("Session token", sessionToken)
	// http.Redirect(w, r, "/profile", http.StatusSeeOther)
	w.Write([]byte("Login successful"))
}

// Logout logs out a user - POST request
func Logout(w http.ResponseWriter, r *http.Request) {
	// TODO - delete session from DB
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
