package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

var (
	ErrNoEmailAddress      = status.Errorf(codes.InvalidArgument, "email address cannot be empty")
	ErrInvalidEmailAddress = status.Errorf(codes.InvalidArgument, "invalid email address")

	ErrNoPassword      = status.Errorf(codes.InvalidArgument, "password cannot be empty")
	ErrInvalidPassword = status.Errorf(codes.InvalidArgument, "password must be at least 8 characters long and may contain at least one special character")
)

func ValidateEmail(email string) error {
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	if len(email) == 0 {
		return ErrNoEmailAddress
	}
	if ok, _ := regexp.MatchString(emailRegex, email); !ok {
		return ErrInvalidEmailAddress
	}

	return nil
}

func ValidatePassword(password string) error {
	passwordRegex := `^[a-zA-Z0-9!@#$&()\-.+]{8,}$`
	if len(password) == 0 {
		return ErrNoPassword
	}
	if ok, _ := regexp.MatchString(passwordRegex, password); !ok {
		return ErrInvalidPassword
	}

	return nil
}
