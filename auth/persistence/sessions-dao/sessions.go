package sessionsdao

import (
	"go-apps/auth.com/db"
	"go-apps/auth.com/model"
	"time"
)

func GetSessionByUserID(userID uint) *model.SessionRecord {
	sql := db.DB
	var session *model.SessionRecord
	result := sql.First(&session, "user_id = ?", userID)
	if result.Error != nil {
		return nil
	}
	return session
}

func DeleteSession(userID uint) error {
	sql := db.DB
	session := GetSessionByUserID(userID)
	if session == nil {
		return nil
	}
	result := sql.Delete(&session)
	return result.Error
}

func CreateSession(user *model.UserRecord, sessionToken string, expiresAt time.Time) {
	db.DB.Create(&model.SessionRecord{
		UserID:    user.ID,
		Email:     user.Email,
		Token:     sessionToken,
		ExpiresAt: expiresAt.Unix(),
	})
}

func GetSessionByToken(token string) *model.SessionRecord {
	var sessionRecord *model.SessionRecord
	sql := db.DB
	result := sql.Where("token = ?", token).First(&sessionRecord)
	if result.RowsAffected == 0 {
		return nil
	}
	return sessionRecord
}
