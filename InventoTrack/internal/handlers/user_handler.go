package handlers

import (
	"encoding/json"
	"inventotrack/internal/services"
	"inventotrack/internal/utils"
	"net/http"
)

// CreateUser handles POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		CompanyID int    `json:"company_id"`
		Role      string `json:"role"` // Add role to the request payload
	}

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Set a default role if none is provided
	if userRequest.Role == "" {
		userRequest.Role = "user" // Default role
	}

	// Call the service to create a user
	user, err := services.CreateUser(userRequest.Username, userRequest.Email, userRequest.Password, userRequest.CompanyID, userRequest.Role)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the created user
	utils.RespondWithJSON(w, http.StatusCreated, user)
}
