package model

import "gorm.io/gorm"

type SessionRecord struct {
	gorm.Model
	UserID uint
	Email  string
	// Token     string `gorm:"uniqueIndex"`
	Token     string
	ExpiresAt int64
}

// TableName overrides the table name instead of session_record it will be sessions
func (u SessionRecord) TableName() string {
	return "sessions"
}
