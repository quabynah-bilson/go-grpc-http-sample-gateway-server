package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

func ValidateEmail(email string) error {
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	if len(email) == 0 {
		return status.Errorf(codes.InvalidArgument, "email address cannot be empty")
	}
	if ok, err := regexp.MatchString(emailRegex, email); err != nil {
		return status.Errorf(codes.Internal, "there was an error validating the email address")
	} else if !ok {
		return status.Errorf(codes.InvalidArgument, "invalid email address")
	}

	return nil
}

func ValidatePassword(password string) error {
	passwordRegex := `^[a-zA-Z0-9!@#$&()\-.+]{8,}$`
	if len(password) == 0 {
		return status.Errorf(codes.InvalidArgument, "password cannot be empty")
	}
	if ok, err := regexp.MatchString(passwordRegex, password); err != nil {
		return status.Errorf(codes.Internal, "there was an error validating the password")
	} else if !ok {
		return status.Errorf(codes.InvalidArgument,
			"invalid password. password must be at least 8 characters long and may contain at least one of the following: !@#$&()-.+")
	}

	return nil
}
