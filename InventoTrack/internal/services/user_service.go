package services

import (
	"errors"
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser creates a new user and checks for duplicates
func CreateUser(username, email, password string, companyID int) (*models.User, error) {
	// Validate inputs
	if username == "" || email == "" || password == "" || companyID <= 0 {
		return nil, errors.New("invalid user data")
	}

	// Check if the username already exists
	existingUser, _ := repositories.GetUserByUsername(username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create the user object
	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		CompanyID:    companyID,
	}

	// Save to database
	err = repositories.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID fetches a user by their ID
func GetUserByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	user, err := repositories.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func DoesUsernameExist(username string) (bool, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return user != nil, nil
}
