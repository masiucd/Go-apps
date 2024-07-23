package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

var DB *sql.DB

func getAllUsers() []User {
	rows, err := DB.Query(`SELECT * FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		fmt.Println("Row")
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)

		}
		users = append(users, user)
	}
	return users
}

func getUserByName(name string) User {
	var user User
	sqlStatement := fmt.Sprintf(`SELECT * FROM users WHERE name like  '%%%s%%' `, name)
	row, err := DB.Query(sqlStatement, name)
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}

	// .Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	return user

}

func main() {
	db := ConnectDB()
	defer db.Close()

	users := getAllUsers()
	fmt.Println(users)

	user := getUserByName("John")
	fmt.Println("user", user)

	// http.HandleFunc("/tasks", tasksHandler)
	// http.HandleFunc("/tasks/", taskHandler)

	fmt.Println("Server is running on port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal(err)
	}
	createTables(db)
	DB = db
	return DB
}

func createTables(db *sql.DB) {
	createTaskTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"email" TEXT,
		"password" TEXT,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(createTaskTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Create users table success")

	createTaskTableSQL = `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"completed" INTEGER DEFAULT 0,
		"user_id" INTEGER,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err = db.Exec(createTaskTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Create tasks table success")
}
