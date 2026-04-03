package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
}
