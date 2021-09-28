package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey"`
	Email        string         `gorm:"type:varchar;uniqueIndex"`
	Password     string         `gorm:"type:varchar"`
	RefreshToken sql.NullString `gorm:"type:text"`
	ExpiresIn    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 4)
	u.Password = string(bytes)
	return
}
