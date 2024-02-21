package main

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func userById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("User with id: " + id))
}

var DB *sql.DB

func main() {
	// db, err := sql.Open("sqlite3", "./db.db")
	// if err != nil {
	// 	panic(err)
	// }
	ConnectDB()

	DB.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")

	http.HandleFunc("GET /", welcome)
	http.HandleFunc("GET /users/{id}", userById)
	http.ListenAndServe(":8080", nil)

	// defer db.Close()

}

func ConnectDB() {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()
}
