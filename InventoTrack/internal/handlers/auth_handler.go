package handlers

import (
	"encoding/json"
	"inventotrack/internal/services"
	"inventotrack/internal/utils"
	"log"
	"net/http"
)

// LoginUser handles POST /login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Printf("Error decoding login request: %v", err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Authenticate the user
	token, err := services.AuthenticateUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// Respond with the token
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
