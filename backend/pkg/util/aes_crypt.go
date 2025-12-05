package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func GetAESDecryptedURLtoJSON(encrypted string) (map[string]interface{}, error) {
	decodedBytes, err := base64.URLEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}
	encrypted = string(decodedBytes)
	decrypted, err := GetAESDecrypted(encrypted)
	if err != nil {
		return nil, err
	}
	var jsonMaps map[string]interface{}
	err = json.Unmarshal(decrypted, &jsonMaps)
	if err != nil {
		return nil, err
	}
	return jsonMaps, nil
}
func GetAESEcryptedURLfromJSON(jsonMaps map[string]interface{}) (string, error) {
	jsonText, err := json.Marshal(jsonMaps)
	if err != nil {
		return "", err
	}
	encripted, err := GetAESEncrypted(string(jsonText))
	if err != nil {
		return "", err
	}
	encripted = base64.URLEncoding.EncodeToString([]byte(encripted))
	return encripted, nil
}

func GetAESDecrypted(encrypted string) ([]byte, error) {
	key := os.Getenv("AES_KEY")
	iv := os.Getenv("AES_KEY_IV")

	if len(key) != 32 || len(iv) != 16 {
		return nil, errors.New("invalid key or iv length")
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		fmt.Println("base64 decode error:", err)
		return nil, err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = PKCS5UnPadding(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func PKCS5UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return nil, errors.New("src is empty")
	}
	unpadding := int(src[length-1])
	if unpadding > length || unpadding == 0 {
		return nil, errors.New("invalid padding size")
	}
	return src[:(length - unpadding)], nil
}

func GetAESEncrypted(plaintext string) (string, error) {
	key := os.Getenv("AES_KEY")
	iv := os.Getenv("AES_KEY_IV")

	if len(key) != 32 || len(iv) != 16 {
		return "", errors.New("invalid key or iv length")
	}

	plainTextBlock := PKCS5Padding([]byte(plaintext), aes.BlockSize)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}
