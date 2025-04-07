package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Items     []CartItem     `gorm:"foreignKey:CartID" json:"items"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CartItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CartID    uint           `json:"cart_id"`
	Cart      Cart           `gorm:"foreignKey:CartID" json:"cart"`
	ProductID uint           `json:"product_id"`
	Product   Product        `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
