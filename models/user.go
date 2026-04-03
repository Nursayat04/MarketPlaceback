package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
}
