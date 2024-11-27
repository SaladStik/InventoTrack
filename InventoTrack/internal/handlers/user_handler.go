package handlers

import (
	"encoding/json"
	"inventotrack/internal/services"
	"inventotrack/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUser handles POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to create a user")

	var userRequest struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		CompanyID int    `json:"company_id"`
	}

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		log.Printf("Error decoding request body: %v", err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	log.Printf("Parsed user request: %+v", userRequest)

	// Call the service to create a user
	user, err := services.CreateUser(userRequest.Username, userRequest.Email, userRequest.Password, userRequest.CompanyID)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("User created successfully: %+v", user)

	// Respond with the created user
	utils.RespondWithJSON(w, http.StatusCreated, user)
}

// GetUser handles GET /users/{id}
func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get a user")

	// Extract the "id" parameter from the URL
	vars := mux.Vars(r) // Extract path parameters
	idStr := vars["id"] // Get the "id" parameter
	log.Printf("Extracted ID from URL: %s", idStr)

	id, err := strconv.Atoi(idStr) // Convert string to int
	if err != nil {
		log.Printf("Invalid user ID: %s", idStr)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Call the service to get the user
	user, err := services.GetUserByID(id)
	if err != nil {
		log.Printf("Error fetching user with ID %d: %v", id, err)
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	log.Printf("User fetched successfully: %+v", user)

	// Respond with the user data
	utils.RespondWithJSON(w, http.StatusOK, user)
}

// CheckUsernameAvailability handles GET /check-username?username={username}
func CheckUsernameAvailability(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to check username availability")

	username := r.URL.Query().Get("username")
	if username == "" {
		log.Println("Username not provided in query parameters")
		utils.RespondWithError(w, http.StatusBadRequest, "Username is required")
		return
	}

	log.Printf("Checking availability for username: %s", username)

	exists, err := services.DoesUsernameExist(username)
	if err != nil {
		log.Printf("Error checking username availability: %v", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if exists {
		log.Printf("Username '%s' is already taken", username)
		utils.RespondWithJSON(w, http.StatusOK, map[string]bool{"available": false})
	} else {
		log.Printf("Username '%s' is available", username)
		utils.RespondWithJSON(w, http.StatusOK, map[string]bool{"available": true})
	}
}
