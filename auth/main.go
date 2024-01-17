package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port = ":8080"

type User struct {
	Name string
	Age  int
}

type UserPageData struct {
	PageTitle string
	Users     []User
}

func users(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/users.html"))
	data := UserPageData{
		PageTitle: "Users Page",
		Users: []User{
			{Name: "John", Age: 25},
			{Name: "Jane", Age: 30},
		},
	}
	tmpl.Execute(w, data)
}

func main() {
	fmt.Println("Server running on port", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.HandleFunc(("/users"), users)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
