package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tsaqiffatih/product-service/config"
	"github.com/tsaqiffatih/product-service/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db := config.InitDB()
	log.Println("Database Connected")

	// Setup router
	r := mux.NewRouter()
	routes.RegisterProductRoutes(r, db)

	port := os.Getenv("PRODUCT_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
