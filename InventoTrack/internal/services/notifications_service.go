package services

import (
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"
)

// GetNotifications fetches all notifications for a user
func GetNotifications(userID int) ([]models.Notification, error) {
	return repositories.GetNotifications(userID)
}

// MarkNotificationAsRead updates a notification's `is_read` status
func MarkNotificationAsRead(notificationID string) error {
	return repositories.MarkNotificationAsRead(notificationID)
}

// CreateNotification creates a new notification (for internal use)
func CreateNotification(userID int, message string) error {
	notification := models.Notification{
		UserID:  userID,
		Message: message,
	}
	return repositories.CreateNotification(notification)
}
