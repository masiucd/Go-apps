package routes

import (
	authmanager "go-apps/auth.com/biz/auth-manager"
	"go-apps/auth.com/lib"

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

	user := authmanager.UserByEmail(email)
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

	session := authmanager.GetSessionByUserId(user.ID)
	if session != nil {
		authmanager.DeleteSession(user.ID)
	}

	token, expiresAt := generateSessionToken()
	authmanager.CreateSession(user, token, expiresAt)

	// store session in cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: expiresAt,
	})

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

// Logout logs out a user - POST request
func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	session := authmanager.GetSessionByToken(c.Value)
	if session != nil {
		err := authmanager.DeleteSession(session.UserID)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   "",
				Expires: time.Now().Add(-time.Hour),
			})
		}
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func generateSessionToken() (string, time.Time) {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(time.Minute * 60) // 1 hour
	return sessionToken, expiresAt
}
