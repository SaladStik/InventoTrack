package services

import (
	"errors"
	"inventotrack/internal/models"
	"inventotrack/internal/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser creates a new user and saves it in the database
func CreateUser(username, email, password string, companyID int, role string) (*models.User, error) {
	// Validate inputs
	if username == "" || email == "" || password == "" {
		return nil, errors.New("username, email, and password are required")
	}
	if role == "" {
		role = "user" // Default role
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create the user object
	companyIDPointer := &companyID // Convert int to *int
	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		CompanyID:    companyIDPointer,
		Role:         role,
	}

	// Save the user to the database
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
