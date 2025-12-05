package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
)

// EncryptInt encrypts an integer and returns the encrypted string
func EncryptInt(key []byte, number int) (string, error) {
	// Convert int to byte slice
	plaintext := make([]byte, 8) // Enough to hold a 64-bit int
	binary.BigEndian.PutUint64(plaintext, uint64(number))

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM (Galois/Counter Mode) for AES
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate a random nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the plaintext
	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)

	// Combine nonce and ciphertext
	combined := append(nonce, ciphertext...)

	// Encode to Base64 for easier handling
	return base64.StdEncoding.EncodeToString(combined), nil
}

// DecryptInt decrypts an encrypted string back into an integer
func DecryptInt(key []byte, encrypted string) (int, error) {
	// Decode Base64
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return 0, err
	}

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	// Create GCM (Galois/Counter Mode) for AES
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return 0, err
	}

	// Split nonce and ciphertext
	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return 0, errors.New("invalid data length")
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt the ciphertext
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return 0, err
	}

	// Convert plaintext back to integer
	number := binary.BigEndian.Uint64(plaintext)
	return int(number), nil
}
