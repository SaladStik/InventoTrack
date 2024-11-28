package repositories

import (
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
)

// AddFeedback saves feedback to the database
func AddFeedback(feedback models.Feedback) error {
	result := extensions.DB.Create(&feedback)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllFeedback retrieves all feedback entries from the database
func GetAllFeedback() ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	result := extensions.DB.Find(&feedbacks)
	if result.Error != nil {
		return nil, result.Error
	}
	return feedbacks, nil
}
