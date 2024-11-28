package repositories

import (
	"inventotrack/internal/extensions"
	"inventotrack/internal/models"
)

// CreateUser saves a new user to the database
func CreateUser(user *models.User) error {
	result := extensions.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(id int) (*models.User, error) {
	var user models.User
	result := extensions.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by their email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := extensions.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByUsername fetches a user by their username
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := extensions.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
