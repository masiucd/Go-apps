package model

import "gorm.io/gorm"

type UserRecord struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// TableName overrides the table name instead of user_records it will be users
func (u UserRecord) TableName() string {
	return "users"
}
