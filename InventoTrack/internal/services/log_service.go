package services

import (
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"
)

// CreateLogEntry creates a log entry for an action
func CreateLogEntry(userID int, action, entity string, entityID int, details map[string]interface{}) error {
	return repositories.CreateLog(userID, entity, entityID, action, details)
}

// FetchAllLogs fetches all logs in the system
func FetchAllLogs() ([]models.Log, error) {
	return repositories.GetLogs()
}

// FetchLogsByUser fetches logs for a specific user
func FetchLogsByUser(userID int) ([]models.Log, error) {
	return repositories.GetLogsByUser(userID)
}

// FetchLogsByEntity fetches logs for a specific entity
func FetchLogsByEntity(entity string, entityID int) ([]models.Log, error) {
	return repositories.GetLogsByEntity(entity, entityID)
}
