package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go-apps/auth.com/db"

	_ "github.com/mattn/go-sqlite3"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index.html").ParseFiles("static/index.html")
	data := struct {
		Title       string
		Description string
	}{
		Title:       "Auth app with Go",
		Description: "Simple auth app using Go and sql-lite",
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template and write the output to the ResponseWriter
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

type User struct {
	ID       int
	Username string
	Email    string
}

func userById(w http.ResponseWriter, r *http.Request) {
	sql := db.DB
	id := r.PathValue("id")
	stmt, err := sql.Prepare("select u.id, u.username, u.email from users u where u.id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var user User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		templ, err := template.New("error.html").ParseFiles("static/error.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		templ.ExecuteTemplate(w, "error.html", "User not found")
		return
	}
	tmpl, err := template.New("user.html").ParseFiles("static/user.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "user.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func users(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10"
	}
	sql := db.DB
	stmt, err := sql.Prepare("select u.id, u.username, u.email from users u limit ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)
	fmt.Println("rows", rows)
	if err != nil {

		log.Fatal(err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		users = append(users, user)
	}

	templ, err := template.New("users.html").ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ.ExecuteTemplate(w, "users.html", users)

}

func main() {
	dir := http.Dir("./static")
	fs := http.FileServer(dir)

	db.ConnectDB()
	defer db.DB.Close()

	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("GET /", welcome)
	mux.HandleFunc("GET /users", users)
	mux.HandleFunc("GET /users/{id}", userById)

	fmt.Println("Server is running on port 9000")
	http.ListenAndServe(":9000", mux)

	// defer db.Close()

}
