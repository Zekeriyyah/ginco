package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zekeriyyah/ginco/internal/models"
)

func GetUserProfile(c *gin.Context) {
	// get the user_id from the request context
	// get user details from db and send it as response

	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	response, err := models.GetUser(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, response)
}
