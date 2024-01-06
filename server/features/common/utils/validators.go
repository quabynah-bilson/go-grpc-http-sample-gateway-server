package utils

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

var (
	ErrNoEmailAddress      = status.Errorf(codes.InvalidArgument, "email address cannot be empty")
	ErrInvalidEmailAddress = status.Errorf(codes.InvalidArgument, "invalid email address")

	ErrNoPassword      = status.Errorf(codes.InvalidArgument, "password cannot be empty")
	ErrInvalidPassword = status.Errorf(codes.InvalidArgument, "password must be at least 8 characters long and may contain at least one special character")

	ErrNoName      = status.Errorf(codes.InvalidArgument, "name cannot be empty")
	ErrInvalidName = status.Errorf(codes.InvalidArgument, "invalid name")

	ErrNoId      = status.Errorf(codes.InvalidArgument, "id cannot be empty")
	ErrInvalidId = status.Errorf(codes.InvalidArgument, "invalid id")
)

func ValidateEmail(email string) error {
	if len(email) == 0 {
		return ErrNoEmailAddress
	}
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
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

func ValidateName(name string) error {
	if len(name) == 0 {
		return ErrNoName
	}

	// must contain only letters and spaces min 2 chars
	nameRegex := `^[a-zA-Z ]{2,}$`
	if ok, _ := regexp.MatchString(nameRegex, name); !ok {
		return ErrInvalidName
	}

	return nil
}

func ValidateId(id string) error {
	if len(id) == 0 {
		return ErrNoId
	}

	if _, err := uuid.Parse(id); err != nil {
		return ErrInvalidId
	}

	return nil
}
