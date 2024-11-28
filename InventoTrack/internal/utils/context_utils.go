package utils

import (
	"context"
)

func GetUserIDFromContext(ctx context.Context) int {
	if ctx == nil {
		return 0
	}
	userID, _ := ctx.Value("userID").(int)
	return userID
}

// GetCompanyIDFromContext retrieves the company ID from the request context
func GetCompanyIDFromContext(ctx context.Context) int {
	if ctx == nil {
		return 0
	}
	companyID, _ := ctx.Value("companyID").(int)
	return companyID
}
