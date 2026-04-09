package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model `json:"-"`
	ID         uint     `json:"id" gorm:"primaryKey"`
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	Stock      int      `json:"stock"`
	UsernameID uint     `json:"user_id"`
	Username   User     `json:"user" gorm:"foreignKey:UserID"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
}
