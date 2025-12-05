package util

import (
	"golang.org/x/crypto/bcrypt"
)

// // HashPassword hashes a plain-text password
// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(bytes), err
// }

// // CheckPasswordHash compares a plain-text password with a hash
//
//	func CheckPasswordHash(password, hash string) error {
//		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//		return err
//	}
//
// Hash password using the bcrypt hashing algorithm
func HashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashedPasswordBytes), nil
}
func CheckPasswordHash(hashedPassword, password string) error {
	// Convert hashed password and input password to byte slices
	hashedPasswordBytes := []byte(hashedPassword)
	inputPasswordBytes := []byte(password)

	// Compare the hashed password with the input password
	return bcrypt.CompareHashAndPassword(hashedPasswordBytes, inputPasswordBytes)
}
