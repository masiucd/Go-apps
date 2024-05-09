package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&UserRecord{})

	DB = db

}
