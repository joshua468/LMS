package main

import (
	"github.com/joshua468/lms-backend/config"
	"github.com/joshua468/lms-backend/routes"
	"log"
	"net/http"
)

func main() {
	// Initialize database
	config.ConnectDatabase()

	// Set up router
	router := routes.SetupRouter()

	// Start the server
	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
