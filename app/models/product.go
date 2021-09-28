package models

import (
	"time"
)

type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar" json:"name"`
	Price     float32   `gorm:"type:decimal(10,2)" json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
