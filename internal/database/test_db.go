package database

import (
	"fmt"
	"log"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

// Setup test database
func TestSetup() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "ginco_user", "ginco_pass", "ginco_test_db", "5040",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to test database:", err)
	}

	TestDB = db
	log.Println("✅ Test Database connected successfully!")
}

// Cleanup test database
func TestTeardown() {
	sqlDB, _ := TestDB.DB()
	sqlDB.Close()
	os.Remove("ginco_test_db")
}

// TestMain runs before all tests
func TestMain(m *testing.M) {
	TestSetup()
	code := m.Run()
	TestTeardown()
	os.Exit(code)
}
