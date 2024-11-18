package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hashing a password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", fmt.Errorf("error to hash password: %w", err)
	}

	return string(bytes), err
}

// Verify hash password
func VerifyHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
