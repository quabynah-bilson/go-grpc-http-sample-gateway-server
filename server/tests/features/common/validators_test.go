package common

import (
	"errors"
	"github.com/eganow/partners/sampler/api/v1/features/common/utils"
	"testing"
)

func TestValidators_ValidateEmailAddress(t *testing.T) {
	cases := []struct {
		name        string
		email       string
		expectedErr error
	}{
		{
			name:        "valid email address",
			email:       "sampler@domain.com",
			expectedErr: nil,
		},
		{
			name:        "no email address",
			email:       "",
			expectedErr: utils.ErrNoEmailAddress,
		},
		{
			name:        "invalid email address",
			email:       "sampler@domain",
			expectedErr: utils.ErrInvalidEmailAddress,
		},
		{
			name:        "improper email address",
			email:       "!2324234@232434.121212",
			expectedErr: utils.ErrInvalidEmailAddress,
		},
		{
			name:        "invalid email address",
			email:       "samplerdomain.com",
			expectedErr: utils.ErrInvalidEmailAddress,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := utils.ValidateEmail(tc.email)

			// Assert
			if !errors.Is(tc.expectedErr, result) {
				t.Errorf("expected: %v, got: %v", tc.expectedErr, result)
			}
		})
	}
}

func TestValidators_ValidatePassword(t *testing.T) {
	cases := []struct {
		name        string
		password    string
		expectedErr error
	}{
		{
			name:        "valid password",
			password:    "Sampler@2024",
			expectedErr: nil,
		},
		{
			name:        "no password",
			password:    "",
			expectedErr: utils.ErrNoPassword,
		},
		{
			name:        "invalid password",
			password:    "sampler",
			expectedErr: utils.ErrInvalidPassword,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := utils.ValidatePassword(tc.password)

			// Assert
			if !errors.Is(tc.expectedErr, result) {
				t.Errorf("expected: %v, got: %v", tc.expectedErr, result)
			}
		})
	}
}

func TestValidators_ValidateName(t *testing.T) {
	cases := []struct {
		name        string
		nameStr     string
		expectedErr error
	}{
		{
			name:        "valid name",
			nameStr:     "Sampler",
			expectedErr: nil,
		},
		{
			name:        "no name",
			nameStr:     "",
			expectedErr: utils.ErrNoName,
		},
		{
			name:        "invalid name",
			nameStr:     "sampler@2024",
			expectedErr: utils.ErrInvalidName,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := utils.ValidateName(tc.nameStr)

			// Assert
			if !errors.Is(tc.expectedErr, result) {
				t.Errorf("expected: %v, got: %v", tc.expectedErr, result)
			}
		})
	}
}

func TestValidators_ValidateId(t *testing.T) {
	cases := []struct {
		name        string
		id          string
		expectedErr error
	}{
		{
			name:        "valid id",
			id:          "D035C330-B399-4E77-9671-1F3AA80DF075",
			expectedErr: nil,
		},
		{
			name:        "no id",
			id:          "",
			expectedErr: utils.ErrNoId,
		},
		{
			name:        "invalid id",
			id:          "1234567890@",
			expectedErr: utils.ErrInvalidId,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := utils.ValidateId(tc.id)

			// Assert
			if !errors.Is(tc.expectedErr, result) {
				t.Errorf("expected: %v, got: %v", tc.expectedErr, result)
			}
		})
	}
}
