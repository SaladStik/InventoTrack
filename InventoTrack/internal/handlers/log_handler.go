package handlers

import (
	"inventotrack/internal/services"
	"inventotrack/internal/utils"
	"net/http"
	"strconv"
)

// GetLogs handles GET /logs
func GetLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := services.FetchAllLogs()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch logs")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, logs)
}

// GetLogsByUser handles GET /logs/user/{userID}
func GetLogsByUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := utils.GetPathParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	logs, err := services.FetchLogsByUser(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch user logs")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, logs)
}

// GetLogsByEntity handles GET /logs/entity/{entity}/{entityID}
func GetLogsByEntity(w http.ResponseWriter, r *http.Request) {
	entity := utils.GetPathParam(r, "entity")
	entityIDStr := utils.GetPathParam(r, "entityID")
	entityID, err := strconv.Atoi(entityIDStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid entity ID")
		return
	}

	logs, err := services.FetchLogsByEntity(entity, entityID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch entity logs")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, logs)
}
