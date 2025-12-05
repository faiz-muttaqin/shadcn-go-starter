package util

import (
	"encoding/base64"
	"fmt"
	"mime"
	"os"
	"path/filepath"
)

func EncodeImageToBase64(filePath string) (string, error) {
	// Baca file gambar
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Tentukan MIME type berdasarkan ekstensi file
	mimeType := mime.TypeByExtension(filepath.Ext(filePath))
	if mimeType == "" {
		return "", fmt.Errorf("unknown file type for %s", filePath)
	}

	// Encode file menjadi base64
	base64Encoding := base64.StdEncoding.EncodeToString(fileBytes)

	// Format string base64 untuk digunakan dalam HTML img src
	base64Image := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Encoding)

	return base64Image, nil
}
