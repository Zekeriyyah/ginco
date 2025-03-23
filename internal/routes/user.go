package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zekeriyyah/ginco/internal/handlers"
	"github.com/zekeriyyah/ginco/internal/middleware"
)

func UserRoutes(r *gin.Engine) {
	// Auth routes
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	// protected routes
	routes := r.Group("/")
	routes.Use(middleware.AuthMiddleware())
	routes.GET("/profile", handlers.GetUserProfile)
}
