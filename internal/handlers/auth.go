package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zekeriyyah/ginco/internal/database"
	"github.com/zekeriyyah/ginco/internal/models"
	"github.com/zekeriyyah/ginco/pkg"
)

func Signup(c *gin.Context) {

	user := &models.User{}

	// bind JSON request to user struct
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// hash password before save
	if err := user.SetPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// save user to database
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection is not initialized"})
		return
	}

	err := database.DB.Create(user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})

}

func Login(c *gin.Context) {
	// get input from request
	// verify email
	// verify password
	// if successful, generate token
	// send response with the token

	var input struct {
		Email    string `json:"email"`
		Password string `json:"-"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	if !user.VerifyPassword(input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error-p": "invalid email or password"})
		return
	}

	tokenStr, err := pkg.GeneratJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}
