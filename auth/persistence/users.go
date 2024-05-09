package persistence

import (
	"go-apps/auth.com/db"
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
