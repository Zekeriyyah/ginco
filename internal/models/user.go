package models

import (
	"fmt"
	"time"

	"github.com/zekeriyyah/ginco/internal/database"
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

// GetUser accepts id (uint) or email (string) to retrieve user from db
func GetUser[T uint | string](value T) (*User, error) {
	// query database for user with id or email, store it in User instance
	// if user exist return user and nil error
	// else return err and empty user struct
	var user User

	switch v := any(value).(type) {
	case uint:
		if err := database.DB.First(&user, v).Error; err != nil {
			return &User{}, err
		}
	case string:
		if err := database.DB.Where("email = ?", v).First(&user).Error; err != nil {
			return &User{}, err
		}
	}

	return &user, nil
}

func GetUsers() ([]User, error) {
	var users []User

	if err := database.DB.Find(&users).Error; err != nil {
		return []User{}, err
	}

	return users, nil
}

func CreateUser(userPtr *User) error {
	if userPtr == nil {
		return fmt.Errorf("user details not provided")
	}

	result := database.DB.Create(userPtr)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateUser(userPtr *User) error {
	if userPtr == nil {
		return fmt.Errorf("user details not provided")
	}

	err := database.DB.Model(userPtr).Updates(*userPtr).Error
	if err != nil {
		return err
	}
	return nil
}
