package main

import (
	"inventotrack/config"
	"inventotrack/internal/extensions"
	"inventotrack/internal/handlers"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()              // Load configurations
	extensions.InitDatabase()        // Initialize the database
	router := handlers.SetupRouter() // Setup API routes

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
