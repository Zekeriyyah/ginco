package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	ProductID uint           `json:"product_id"`
	Product   Product        `gorm:"foreignKey:ProductID" json:"product"`
	Rating    int            `gorm:"not null" json:"rating"` // 1-5 stars
	Comment   string         `json:"comment"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
