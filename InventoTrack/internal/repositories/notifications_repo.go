package repositories

import (
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
)

// GetNotifications retrieves all notifications for a user
func GetNotifications(userID int) ([]models.Notification, error) {
	var notifications []models.Notification
	result := extensions.DB.Where("user_id = ?", userID).Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return notifications, nil
}

// MarkNotificationAsRead sets a notification's `is_read` field to true
func MarkNotificationAsRead(notificationID string) error {
	result := extensions.DB.Model(&models.Notification{}).
		Where("id = ?", notificationID).
		Update("is_read", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateNotification inserts a new notification into the database
func CreateNotification(notification models.Notification) error {
	result := extensions.DB.Create(&notification)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
