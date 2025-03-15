package models

import (
	"time"

	"github.com/zekeriyyah/ginco/pkg"
	_ "gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex; not null" json:"username,omitempty"`
	Email     string    `gorm:"uniqueIndex; not null" json:"email,omitempty"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

func (u *User) SetPassword(password string) error {
	hashed, err := pkg.HashPassword(password)
	if err != nil {
		pkg.Error("failed to hash password", err)
		return err
	}

	u.Password = string(hashed)
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	return pkg.CheckPassword(u.Password, password)
}
