package routes

import (
	"github.com/gorilla/mux"
	"github.com/tsaqiffatih/product-service/handlers"
	"gorm.io/gorm"
)

func RegisterProductRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/product", handlers.GetProducts(db)).Methods("GET")
	r.HandleFunc("/product/{id}", handlers.GetProduct(db)).Methods("GET")
	r.HandleFunc("/product", handlers.CreateProduct(db)).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct(db)).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct(db)).Methods("DELETE")

}
