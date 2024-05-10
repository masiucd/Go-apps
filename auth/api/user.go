package api

import (
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	// Parse form data and get firstname and lastname
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.Form["firstname"])
	fmt.Println(r.Form["lastname"])
	fmt.Println(r.Form["email"])
	fmt.Println(r.Form["password"])
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])

	// var input input.UserInput

	// Create user and

	// Redirect to user profile page
}
