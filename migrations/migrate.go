package migrations

import (
	"log"

	"github.com/zekeriyyah/ginco/internal/database"
	"github.com/zekeriyyah/ginco/internal/models"
)

// RunMigrations applies database schema changes
func Run() {
	db := database.DB

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("✅ Migrations completed successfully!")
}
