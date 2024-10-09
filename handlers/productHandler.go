package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/tsaqiffatih/product-service/models"
	"gorm.io/gorm"
)

var validate *validator.Validate

// Initialize validator instance
func init() {
	validate = validator.New()
}

// Get All Products
func GetProducts(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var products []models.Product
		db.Find(&products)
		json.NewEncoder(w).Encode(products)
	}
}

// Get single product by ID
func GetProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(product)
	}
}

// Create a new product
func CreateProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product
		json.NewDecoder(r.Body).Decode(&product)
		db.Create(&product)
		json.NewEncoder(w).Encode(product)
	}
}

// Update product by ID
func UpdateProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		json.NewDecoder(r.Body).Decode(&product)
		db.Save(&product)
		json.NewEncoder(w).Encode(product)
	}
}

// Delete product by ID
func DeleteProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		db.Delete(&product)
		w.WriteHeader(http.StatusNoContent)
	}
}
