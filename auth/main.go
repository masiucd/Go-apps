package main

import (
	"fmt"
	"net/http"

	"go-apps/auth.com/db"
	routes "go-apps/auth.com/routes/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dir := http.Dir("./static")
	fs := http.FileServer(dir)

	db.ConnectDB()
	defer db.DB.Close()

	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("GET /users", routes.UserById)
	mux.HandleFunc("GET /users/{id}", routes.UserById)

	fmt.Println("Server is running on port 9000")
	http.ListenAndServe(":9000", mux)

}
