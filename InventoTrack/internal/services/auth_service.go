package services

import (
	"errors"
	"inventotrack/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key") // Replace with a strong secret key

// AuthenticateUser verifies username and password, and returns a JWT if valid
func AuthenticateUser(username, password string) (string, error) {
	// Fetch the user from the database
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
	})

	// Sign and return the token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}
