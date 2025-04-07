package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zekeriyyah/ginco/internal/database"
	"github.com/zekeriyyah/ginco/internal/routes"
	"github.com/zekeriyyah/ginco/migrations"
	"github.com/zekeriyyah/ginco/pkg"
)

func main() {
	database.InitDB()
	if os.Getenv("RENDER") == "true" {
		migrations.Run()
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.UserRoutes(r)

	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		pkg.Error("server connection failed", err)
	}

	pkg.Info("Server successfully running on port 8080")
}
