package repositories

import (
	"encoding/json"
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
)

// CreateLog creates a new log entry in the database
func CreateLog(userID int, entity string, entityID int, action string, details map[string]interface{}) error {
	// Convert details to JSON
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		return err
	}

	// Create the log entry
	log := &models.Log{
		UserID:   userID,
		Action:   action,
		Entity:   entity,
		EntityID: entityID,
		Details:  string(detailsJSON), // Store as JSON string
	}

	// Save to the database
	result := extensions.DB.Create(log)
	return result.Error
}

// GetLogs retrieves all logs from the database
func GetLogs() ([]models.Log, error) {
	var logs []models.Log
	result := extensions.DB.Order("created_at DESC").Find(&logs)
	return logs, result.Error
}

// GetLogsByUser retrieves logs by a specific user
func GetLogsByUser(userID int) ([]models.Log, error) {
	var logs []models.Log
	result := extensions.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&logs)
	return logs, result.Error
}

// GetLogsByEntity retrieves logs related to a specific entity and entity ID
func GetLogsByEntity(entity string, entityID int) ([]models.Log, error) {
	var logs []models.Log
	result := extensions.DB.Where("entity = ? AND entity_id = ?", entity, entityID).Order("created_at DESC").Find(&logs)
	return logs, result.Error
}
