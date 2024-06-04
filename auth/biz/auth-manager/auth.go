package authmanager

import (
	"go-apps/auth.com/model"
	sessionsdao "go-apps/auth.com/persistence/sessions-dao"
	usersdao "go-apps/auth.com/persistence/users-dao"
	"time"
)

func UserByEmail(email string) *model.UserRecord {
	return usersdao.UserByEmail(email)
}

func GetSessionByUserId(userID uint) *model.SessionRecord {
	return sessionsdao.GetSessionByUserID(userID)
}

func GetSessionByToken(token string) *model.SessionRecord {
	return sessionsdao.GetSessionByToken(token)
}

func DeleteSession(userID uint) error {
	return sessionsdao.DeleteSession(userID)
}

func CreateSession(user *model.UserRecord, token string, expiresAt time.Time) {
	sessionsdao.CreateSession(user, token, expiresAt)
}
