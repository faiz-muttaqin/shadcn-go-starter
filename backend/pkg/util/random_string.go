package util

import (
	"math/rand"
	"time"
)

func GenerateRandomString(charNum int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	token := make([]byte, charNum)

	for i := range token {
		token[i] = charset[r.Intn(len(charset))]
	}

	return string(token)
}
func GenerateRandomNumberString(charNum int) string {
	const charset = "0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	token := make([]byte, charNum)

	for i := range token {
		token[i] = charset[r.Intn(len(charset))]
	}

	return string(token)
}
func GenerateRandomHexaString(charNum int) string {
	const charset = "abcdef0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	token := make([]byte, charNum)

	for i := range token {
		token[i] = charset[r.Intn(len(charset))]
	}

	return string(token)
}
func GenerateRandomStringLowerCase(charNum int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	token := make([]byte, charNum)

	for i := range token {
		token[i] = charset[r.Intn(len(charset))]
	}

	return string(token)
}
func GenerateRandomStringUpperCase(charNum int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	token := make([]byte, charNum)

	for i := range token {
		token[i] = charset[r.Intn(len(charset))]
	}

	return string(token)
}
