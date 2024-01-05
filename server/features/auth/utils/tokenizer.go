package utils

import "github.com/google/uuid"

// GenerateAccessToken generates a new access token (for now it's a UUID v4).
func GenerateAccessToken(_ string) string {
	return uuid.NewString()
}

// GenerateRefreshToken generates a new refresh token (for now it's a UUID v4).
func GenerateRefreshToken(_ string) string {
	return uuid.NewString()
}
