package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	FileName string
	Blob     []byte
}

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Attachment{})
	DB = db
}
