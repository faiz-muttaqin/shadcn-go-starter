package util

import (
	"bytes"
	"io"
	"mime/multipart"
)

func IsValidImage(file multipart.File) (bool, string) {
	// Read the first few bytes of the file
	header := make([]byte, 8)
	if _, err := file.Read(header); err != nil {
		return false, ""
	}

	// Reset the file reader after reading the header
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return false, ""
	}

	if IsPNG(header) {
		return true, "PNG"
	} else if IsJPG(header) {
		return true, "JPG"
	}

	return false, ""
}

func IsPNG(header []byte) bool {
	pngSignature := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	return bytes.Equal(header, pngSignature)
}

func IsJPG(header []byte) bool {
	jpgSignature := []byte{0xFF, 0xD8}
	return bytes.HasPrefix(header, jpgSignature)
}
