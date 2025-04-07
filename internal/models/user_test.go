package models

import (
	"testing"

	"github.com/zekeriyyah/ginco/internal/database"

	"github.com/stretchr/testify/assert"
)

// Test user creation
func TestCreateUser(t *testing.T) {
	// db := database.TestDB

	user := User{
		Username: "testuser",
		Email:    "test@email.com",
		Password: "securepassword",
	}

	err := database.TestDB.Create(&user).Error
	assert.Nil(t, err)
	assert.NotZero(t, user.ID)
}

// Test unique email constraint
func TestUniqueEmail(t *testing.T) {
	db := database.TestDB

	user1 := &User{Username: "user1", Email: "duplicate@email.com", Password: "pass123"}
	user2 := &User{Username: "user2", Email: "duplicate@email.com", Password: "pass456"}

	_ = db.Create(user1)
	err := db.Create(user2).Error

	assert.NotNil(t, err) // Expect error because email is unique
}

// Test password hashing
func TestPasswordHashing(t *testing.T) {
	user := User{Password: "mypassword"}
	err := user.SetPassword(user.Password)

	assert.Nil(t, err)
	assert.NotEqual(t, "mypassword", user.Password) // Should be hashed
}

// Test password verification
func TestPasswordVerification(t *testing.T) {
	user := User{}
	_ = user.SetPassword("mypassword")

	assert.True(t, user.VerifyPassword("mypassword")) // Correct password
	assert.False(t, user.VerifyPassword("wrongpass")) // Wrong password
}
