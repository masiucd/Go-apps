package main

import (
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
		log.Fatal(err)
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
	// w.Write([]byte("User with id: " + fmt.Sprintf("%d", user.ID) + " username = " + user.Username))

}

func main() {
	dir := http.Dir("./static")
	fs := http.FileServer(dir)
	db.ConnectDB()

	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("GET /", welcome)
	mux.HandleFunc("GET /users/{id}", userById)
	http.ListenAndServe(":9000", mux)

	// defer db.Close()

}
