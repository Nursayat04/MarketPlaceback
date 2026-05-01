package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primaryKey"`
	ProductID  uint   `json:"product_id"`
	UserID     uint   `json:"user_id"`
	Username   string `json:"username"`
	Rating     int    `json:"rating"` // 1–5
	Comment    string `json:"comment"`
}
