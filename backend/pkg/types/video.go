package types

import (
	"path/filepath"
	"strings"
)

type Video string

func (v Video) String() string  { return string(v) }
func (v Video) Kind() FieldType { return FieldVideo }

// Ext returns the file extension (.mp4, .mov, .avi, etc)
func (v Video) Ext() string {
	ext := filepath.Ext(string(v))
	return strings.ToLower(ext)
}

// IsVideo checks if file extension is a valid video type
func (v Video) IsVideo() bool {
	switch v.Ext() {
	case ".mp4", ".mov", ".avi", ".mkv", ".webm", ".flv", ".wmv", ".mpeg", ".mpg":
		return true
	}
	return false
}
func (v Video) MimeType() string {
	ext := v.Ext()
	switch ext {
	case ".mp4":
		return "video/mp4"
	case ".mov":
		return "video/quicktime"
	case ".webm":
		return "video/webm"
	case ".avi":
		return "video/x-msvideo"
	case ".mkv":
		return "video/x-matroska"
	}
	return "application/octet-stream"
}
