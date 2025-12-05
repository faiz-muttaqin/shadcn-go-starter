package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func SignatureGenerator(message []byte, key []byte) string {
	mac := hmac.New(sha256.New, key)

	// Write the message to the HMAC
	mac.Write(message)

	// Get the final HMAC result
	expectedMAC := mac.Sum(nil)

	// Encode the HMAC result to Base64
	expectedMACBase64 := base64.StdEncoding.EncodeToString(expectedMAC)
	return expectedMACBase64

}
