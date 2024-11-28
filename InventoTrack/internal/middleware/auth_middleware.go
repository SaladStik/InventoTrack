package middleware

import (
	"context"
	"errors"
	"inventotrack/internal/utils"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key") // Replace with your actual secret key

// AuthMiddleware validates the JWT and protects routes
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header missing")
			return
		}

		// Ensure "Bearer" format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		// Extract the token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token signing method")
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			log.Printf("Invalid token: %v", err)
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		// Extract userID and role from claims with additional checks
		userID, userIDOk := claims["userID"].(float64)
		role, roleOk := claims["role"].(string)
		if !userIDOk || !roleOk {
			log.Printf("Missing claims in token: userID=%v, role=%v", userIDOk, roleOk)
			utils.RespondWithError(w, http.StatusUnauthorized, "User ID or role not found in token")
			return
		}

		// Add userID and role to the request context
		ctx := context.WithValue(r.Context(), "userID", int(userID))
		ctx = context.WithValue(ctx, "role", role)

		// Call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RoleMiddleware ensures the user has the required role
func RoleMiddleware(requiredRole string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, err := GetUserRoleFromContext(r.Context())
		if err != nil || role != requiredRole {
			utils.RespondWithError(w, http.StatusForbidden, "You do not have the required permissions")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// GetUserIDFromContext retrieves the user ID from the request context
func GetUserIDFromContext(ctx context.Context) (int, error) {
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}
	return userID, nil
}

// GetUserRoleFromContext retrieves the user role from the request context
func GetUserRoleFromContext(ctx context.Context) (string, error) {
	role, ok := ctx.Value("role").(string)
	if !ok {
		return "", errors.New("user role not found in context")
	}
	return role, nil
}
