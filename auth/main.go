package main

import "net/http"

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func userById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("User with id: " + id))
}

func main() {

	http.HandleFunc("GET /", welcome)
	http.HandleFunc("GET /users/{id}", userById)

	http.ListenAndServe(":8080", nil)
}
