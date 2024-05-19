package persistence

import (
	"fmt"
	"go-apps/auth.com/db"
	"go-apps/auth.com/input"
	"go-apps/auth.com/model"
)

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

// DoesUserExist
func DoesUserExist(email string) bool {
	sql := db.DB
	var user model.UserRecord
	fmt.Println("email", email)
	result := sql.Where("email = ?", email).First(&user)
	fmt.Println("result.RowsAffected", result.RowsAffected)
	fmt.Println("user", user.FirstName)

	return result.RowsAffected > 0

}
