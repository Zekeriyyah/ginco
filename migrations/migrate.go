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

// Run migrations related to test database
func TestRun() {
	db := database.TestDB

	// Auto-migrate test tables
	err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	if err != nil {
		log.Fatal("❌ Failed to migrate test database:", err)
	}

	log.Println("✅ Test Migrations completed successfully!")
}
