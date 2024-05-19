package model

import "gorm.io/gorm"

type UserRecord struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u UserRecord) TableName() string {
	return "users"
}
