package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// VerifyPassword verifies password hash against plain password
func VerifyPassword(hashedPassword string, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))

	return err == nil
}
