package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `json:"user_id"`
	User       User           `gorm:"foreignKey:UserID" json:"user"`
	TotalPrice float64        `gorm:"not null" json:"total_price"`
	Status     string         `gorm:"not null;default:'pending'" json:"status"` // pending, completed, canceled
	Items      []OrderItem    `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
