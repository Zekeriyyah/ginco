package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrderID   uint           `gorm:"unique;not null" json:"order_id"`
	Order     Order          `gorm:"foreignKey:OrderID" json:"order"`
	Amount    float64        `gorm:"not null" json:"amount"`
	Status    string         `gorm:"not null;default:'pending'" json:"status"` // pending, successful, failed
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
