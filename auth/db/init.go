package db

import "database/sql"

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		panic(err)
	}
	DB = db
	// defer db.Close()
}
