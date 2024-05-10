package routes

import (
	routes "go-apps/auth.com/routes/user"
	"net/http"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/{id}", routes.UserById)
	mux.HandleFunc("GET /users", routes.UserById)
	mux.HandleFunc("GET /signup", routes.SignUp)
	mux.HandleFunc("GET /login", routes.Login)
	mux.HandleFunc("POST /create-user", routes.CreateUser)

	return mux
}
