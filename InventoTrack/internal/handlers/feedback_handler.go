package handlers

import (
	"encoding/json"
	"inventotrack/internal/services"
	"inventotrack/internal/utils"
	"net/http"
)

// SubmitFeedback handles POST /feedback
func SubmitFeedback(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetUserIDFromContext(r.Context())
	if userID == 0 {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or missing user ID")
		return
	}

	var request struct {
		Message string `json:"message"`
	}

	// Parse the request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if request.Message == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Feedback message cannot be empty")
		return
	}

	// Call the service to submit feedback
	err := services.SubmitFeedback(userID, request.Message)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Feedback submitted successfully"})
}

// GetFeedback handles GET /feedback (optional admin-only endpoint)
func GetFeedback(w http.ResponseWriter, r *http.Request) {
	feedbacks, err := services.GetAllFeedback()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, feedbacks)
}
