package types

import (
	"path/filepath"
	"strings"
)

type Document string

func (d Document) String() string  { return string(d) }
func (d Document) Kind() FieldType { return FieldDocument }

func (d Document) Ext() string {
	return strings.ToLower(filepath.Ext(string(d)))
}

func (d Document) IsDocument() bool {
	switch d.Ext() {
	case ".pdf",
		".doc", ".docx",
		".xls", ".xlsx",
		".ppt", ".pptx",
		".txt", ".rtf":
		return true
	}
	return false
}
