package model

type UserRecord struct {
	BaseModel
	FirstName string
	LastName  string
	Email     string
	Password  string
	// DeletedAt gorm.DeletedAt `gorm:"-"` // ignore this field
}

// TableName overrides the table name instead of user_records it will be users
func (u UserRecord) TableName() string {
	return "users"
}
