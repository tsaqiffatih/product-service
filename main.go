package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/tsaqiffatih/product-service/config"
	"github.com/tsaqiffatih/product-service/routes"
	"github.com/tsaqiffatih/product-service/utils"
)

func main() {
	err := utils.LoadPrivateKey("private.key")
	if err != nil {
		log.Fatalf("Gagal memuat private key: %v", err)
	}

	err = utils.LoadPublicKey("public.key")
	if err != nil {
		log.Fatalf("Gagal memuat public key: %v", err)
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
