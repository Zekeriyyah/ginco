package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrderID   uint           `json:"order_id"`
	Order     Order          `gorm:"foreignKey:OrderID" json:"order"`
	ProductID uint           `json:"product_id"`
	Product   Product        `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	Price     float64        `gorm:"not null" json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
