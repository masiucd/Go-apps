package main

import (
	"log"
	"net/http"

	"go-apps/auth.com/db"

	_ "github.com/mattn/go-sqlite3"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func userById(w http.ResponseWriter, r *http.Request) {
	sql := db.DB
	id := r.PathValue("id")
	stmt, err := sql.Prepare("select u.username from users u where u.id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	var username string
	err = stmt.QueryRow("1").Scan(&username)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("User with id: " + id + " username = " + username))

}

func main() {
	db.ConnectDB()
	// db, err := sql.Open("sqlite3", "./db.db")
	// if err != nil {
	// 	panic(err)
	// }

	http.HandleFunc("GET /", welcome)
	http.HandleFunc("GET /users/{id}", userById)
	http.ListenAndServe(":9000", nil)

	// defer db.Close()

}
