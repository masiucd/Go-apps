package main

import (
	"fmt"
	"net/http"

	"go-apps/auth.com/api"
	"go-apps/auth.com/db"
	routes "go-apps/auth.com/routes/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// dir := http.Dir("./static")
	// fs := http.FileServer(dir)

	db.ConnectDB()

	mux := http.NewServeMux()
	// mux.Handle("/", fs)
	mux.HandleFunc("GET /users/{id}", routes.UserById)
	mux.HandleFunc("GET /users", routes.UserById)
	mux.HandleFunc("GET /signup", routes.SignUp)
	mux.HandleFunc("POST /create-user", api.CreateUser)

	fmt.Println("Server is running on port 4000")
	http.ListenAndServe(":4000", mux)

}
