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
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		search := r.URL.Query().Get("search")

		if limit == 0 {
			limit = 10 // Default limit
		}

		var products []models.Product
		query := db.Limit(limit).Offset(offset)

		if search != "" {
			query = query.Where("name LIKE ?", "%"+search+"%")
		}

		query.Find(&products)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "success",
			"message": "Get Products successfully",
			"data":    products,
		})
	}
}

// Get single product by ID
func GetProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "fail",
				"error": map[string]interface{}{
					"code":    http.StatusNotFound,
					"message": "Product not found",
				},
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "success",
			"message": "Get Product successfully",
			"data":    product,
		})
	}
}

// Create a new product
func CreateProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product

		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "fail",
				"error": map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "Invalid request payload",
				},
			})
			return
		}

		// validation input
		err = validate.Struct(product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "fail",
				"error": map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": err.Error(),
				},
			})
			return
		}

		if err := db.Create(&product).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "fail",
				"error": map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "Failed to create product",
				},
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "success",
			"message": "Create Product successfully",
			"data":    product,
		})
	}
}

// Update product by ID
func UpdateProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "fail",
				"error": map[string]interface{}{
					"code":    http.StatusNotFound,
					"message": "Product not found",
				},
			})
			return
		}

		json.NewDecoder(r.Body).Decode(&product)
		db.Save(&product)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "success",
			"message": "Update Product successfully",
			"data":    product,
		})
	}
}

// Delete product by ID
func DeleteProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "fail",
				"error": map[string]interface{}{
					"code":    http.StatusNotFound,
					"message": "Product not found",
				},
			})
			return
		}

		// soft delete
		if err := db.Delete(&product).Error; err != nil {
			// http.Error(w, "Failed to delete product", http.StatusInternalServerError)
			// return
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "fail",
				"error": map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "Failed to delete product",
				},
			})
			return
		}

		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "success",
			"message": "Product deleted successfully",
		})
	}
}
