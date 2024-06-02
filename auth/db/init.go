package db

import (
	"go-apps/auth.com/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.UserRecord{})
	db.AutoMigrate(&model.SessionRecord{})

	DB = db

}
