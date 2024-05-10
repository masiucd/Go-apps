package routes

import (
	"go-apps/auth.com/api"
	routes "go-apps/auth.com/routes/user"
	"net/http"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.Handle("/", fs)

	mux.HandleFunc("GET /users/{id}", routes.UserById)
	mux.HandleFunc("GET /users", routes.UserById)
	mux.HandleFunc("GET /signup", routes.SignUp)
	mux.HandleFunc("POST /create-user", api.CreateUser)

	return mux
}
