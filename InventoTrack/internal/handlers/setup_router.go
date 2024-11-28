package handlers

import (
	"inventotrack/internal/middleware"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the router and configures all routes
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Public routes (no authentication required)
	router.HandleFunc("/login", LoginUser).Methods("POST")                             // Login user
	router.HandleFunc("/companies", CreateCompany).Methods("POST")                     // Create a company
	router.HandleFunc("/companies-with-owner", CreateCompanyWithOwner).Methods("POST") // Create a company with an owner

	// Protected routes (require authentication)
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware) // Apply authentication middleware to all protected routes

	// User routes
	protected.HandleFunc("/users", CreateUser).Methods("POST") // Create user

	// Inventory routes
	protected.HandleFunc("/inventory", AddInventory).Methods("POST")                      // Add inventory
	protected.HandleFunc("/inventory/{id}", GetInventory).Methods("GET")                  // Get inventory by ID
	protected.HandleFunc("/inventory/{id}/archive", ArchiveInventory).Methods("POST")     // Archive inventory
	protected.HandleFunc("/inventory/{id}/unarchive", UnarchiveInventory).Methods("POST") // Unarchive inventory

	// Custom fields routes
	protected.HandleFunc("/custom-fields", AddCustomField).Methods("POST")           // Add a custom field
	protected.HandleFunc("/custom-fields/{id}", RemoveCustomField).Methods("DELETE") // Remove a custom field

	// Recycle Bin routes
	protected.HandleFunc("/recycle-bin", GetRecycleBin).Methods("GET")             // Get recycle bin entries
	protected.HandleFunc("/recycle-bin/{id}", PermanentlyDelete).Methods("DELETE") // Permanently delete a recycle bin entry

	// Notifications routes
	protected.HandleFunc("/notifications", GetNotifications).Methods("GET") // Get notifications

	// Feedback routes
	protected.HandleFunc("/feedback", SubmitFeedback).Methods("POST") // Submit feedback

	// Logs routes
	protected.HandleFunc("/logs", GetLogs).Methods("GET")                                    // Get all logs
	protected.HandleFunc("/logs/user/{userID}", GetLogsByUser).Methods("GET")                // Get logs by user ID
	protected.HandleFunc("/logs/entity/{entity}/{entityID}", GetLogsByEntity).Methods("GET") // Get logs for a specific entity

	// Apply global middleware (e.g., logging)
	router.Use(middleware.LoggingMiddleware)

	return router
}
