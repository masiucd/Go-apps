package model

type SessionRecord struct {
	BaseModel
	UserID uint
	Email  string
	// Token     string `gorm:"uniqueIndex"`
	Token     string
	ExpiresAt int64
	// DeletedAt gorm.DeletedAt `gorm:"-"` // ignore this field
}

// TableName overrides the table name instead of session_record it will be sessions
func (u SessionRecord) TableName() string {
	return "sessions"
}
