package services

import (
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"
)

// SubmitFeedback saves user feedback to the database
func SubmitFeedback(userID int, message string) error {
	feedback := models.Feedback{
		UserID:  userID,
		Message: message,
	}
	return repositories.AddFeedback(feedback)
}

// GetAllFeedback fetches all feedback from the database (admin-only feature)
func GetAllFeedback() ([]models.Feedback, error) {
	return repositories.GetAllFeedback()
}
