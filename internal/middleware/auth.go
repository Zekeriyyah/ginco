package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zekeriyyah/ginco/pkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		// access the authHeader for the token
		tokenHeader = strings.TrimSpace(tokenHeader)
		tokenSlice := strings.SplitN(tokenHeader, " ", 2)

		if len(tokenSlice) != 2 || tokenSlice[1] == "" {
			pkg.Info("Slicing the bearer token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.Abort()
			return
		}

		token := tokenSlice[1]
		pkg.Info(token)

		claims, err := pkg.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
