package handlers

import (
	"inventotrack/internal/services"
	"inventotrack/internal/utils"
	"net/http"
)

// GetNotifications handles GET /notifications
func GetNotifications(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetUserIDFromContext(r.Context())
	if userID == 0 {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or missing user ID")
		return
	}

	notifications, err := services.GetNotifications(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, notifications)
}

// MarkNotificationAsRead handles POST /notifications/{id}/read
func MarkNotificationAsRead(w http.ResponseWriter, r *http.Request) {
	notificationID := utils.GetPathParam(r, "id")
	if notificationID == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Notification ID is required")
		return
	}

	err := services.MarkNotificationAsRead(notificationID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Notification marked as read"})
}
