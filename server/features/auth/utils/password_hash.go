package utils

import "golang.org/x/crypto/bcrypt"

// EncryptPassword encrypts the password using bcrypt.
func EncryptPassword(password string) (*string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	encryptedPassword := string(bcryptPassword)
	return &encryptedPassword, nil
}

// ComparePasswords compares the password with the encrypted password.
func ComparePasswords(password, encryptedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
}
