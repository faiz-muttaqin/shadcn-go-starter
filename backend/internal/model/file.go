package model

import "io"

type File struct {
	Name     string    // Name of the file
	File     io.Reader // The file content (streamed)
	Size     int64     // Size of the file (optional but useful for efficient transmission)
	MIMEType string    // MIME Type (image/jpeg, application/pdf, etc.)
	FileType string    // The file content type
	FilePath string    // URL if the file is being fetched from a remote location (optional)
}
