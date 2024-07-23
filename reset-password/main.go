package main

import (
	"fmt"
	"log"
	"net/http"

	"go-apps/reset-password.com/persistence"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := persistence.ConnectDB()
	defer db.Close()

	users := persistence.GetAllUsers()
	fmt.Println(users)

	user := persistence.GetUserByName("John")
	fmt.Println("user", user)

	fmt.Println("Server is running on port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
