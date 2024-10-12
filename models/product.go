package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string     `json:"name" gorm:"not null" validate:"required"`
	Price       int        `json:"price" gorm:"not null" validate:"required"`
	Stock       int        `json:"stock" gorm:"default:0" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Categories  []Category `gorm:"many2many:product_categories"`
}

type Category struct {
	gorm.Model
	Name     string    `json:"name" gorm:"not null" validate:"required"`
	Products []Product `gorm:"many2many:product_categories"`
}
