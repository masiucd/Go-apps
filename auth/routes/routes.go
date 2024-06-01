package routes

import (
	"net/http"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", Users)
	mux.HandleFunc("GET /users/{id}", UserById)

	mux.HandleFunc("GET /signup", SignUp)
	mux.HandleFunc("POST /create-user", CreateUser)

	mux.HandleFunc("GET /login", Login)
	mux.HandleFunc("POST /login", LoginUser)

	return mux
}
