package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string     `json:"name" validate:"required"`
	Price       int        `json:"price" validate:"required"`
	Stock       int        `json:"stock" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Categories  []Category `gorm:"many2many:product_categories"`
}

type Category struct {
	gorm.Model
	Name     string    `json:"name" validate:"required"`
	Products []Product `gorm:"many2many:product_categories"`
}
