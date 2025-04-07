package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Price       float64        `gorm:"not null" json:"price"`
	Stock       int            `gorm:"not null;default:0" json:"stock"`
	CategoryID  uint           `json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
