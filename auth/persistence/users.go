package persistence

import (
	"go-apps/auth.com/db"
	"go-apps/auth.com/input"
	"go-apps/auth.com/model"
)

// Users returns a list of users
func Users(limit int) ([]*model.UserRecord, error) {
	sql := db.DB
	var users []*model.UserRecord
	sql.Limit(limit).Find(&users)
	return users, nil
}

func User(id string) *model.UserRecord {
	sql := db.DB
	var user model.UserRecord
	sql.First(&user, id)
	return &user
}

// InsertUser inserts a user into the database
func InsertUser(input input.UserInput) error {
	sql := db.DB
	user := model.UserRecord{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}
	record := sql.Create(&user)
	return record.Error
}

// DoesUserExist checks if a user exists in the database
func DoesUserExist(email string) bool {
	sql := db.DB
	var user model.UserRecord
	result := sql.Where("email = ?", email).First(&user)
	// if RowsAffected is greater than 0, then the user exists
	return result.RowsAffected > 0
}
