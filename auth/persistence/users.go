package persistence

import (
	"go-apps/auth.com/db"
	"go-apps/auth.com/input"
)

func Users(limit int) ([]*db.UserRecord, error) {
	sql := db.DB
	var users []*db.UserRecord
	sql.Limit(limit).Find(&users)
	return users, nil
}

func User(id string) *db.UserRecord {
	sql := db.DB
	var user db.UserRecord
	sql.First(&user, id)
	return &user
}

func InsertUser(input input.UserInput) error {
	sql := db.DB
	user := db.UserRecord{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}
	record := sql.Create(&user)
	return record.Error
}
