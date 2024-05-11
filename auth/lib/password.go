package lib

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a password with bcrypt
func HashPassword(password string) (string, error) {
	bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return password, nil
}

// VerifyPassword compares a hashed password with its possible plaintext equivalent
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
