package db

import (
	"log"
)

type UserRecord struct {
	ID       int
	Username string
	Email    string
}

func Users(limit string) ([]UserRecord, error) {
	sql := DB
	stmt, err := sql.Prepare("select u.id, u.username, u.email from users u limit ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []UserRecord
	for rows.Next() {
		var user UserRecord
		err = rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func User(id string) (*UserRecord, error) {
	sql := DB
	stmt, err := sql.Prepare("select u.id, u.username, u.email from users u where u.id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var user UserRecord
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, err
}
