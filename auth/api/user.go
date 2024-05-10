package api

import (
	"go-apps/auth.com/input"
	"go-apps/auth.com/persistence"
	"net/http"
)

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
