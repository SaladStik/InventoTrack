package handlers

import (
	"inventotrack/internal/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the router and configures all routes
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Add global middleware (e.g., logging)
	router.Use(loggingMiddleware)

	// Public routes (no authentication required)
	public := router.PathPrefix("/").Subrouter()
	public.HandleFunc("/login", LoginUser).Methods("POST")
	public.HandleFunc("/companies", CreateCompany).Methods("POST") // Keep company creation public

	// Protected routes (require authentication)
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware) // Add authentication middleware for protected routes

	// User routes
	protected.HandleFunc("/users", CreateUser).Methods("POST")
	protected.HandleFunc("/users/{id}", GetUser).Methods("GET")

	// Inventory routes
	protected.HandleFunc("/inventory", AddInventory).Methods("POST")
	protected.HandleFunc("/inventory/{id}", GetInventory).Methods("GET")

	return router
}

// loggingMiddleware logs the details of each incoming request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request - Method: %s, Path: %s, RemoteAddr: %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
