package types

import (
	"html/template"
	"path/filepath"
	"strings"
)

type HTML template.HTML

func (i HTML) String() string  { return string(i) }
func (i HTML) Kind() FieldType { return FieldHTML }

// Ext returns the extension of the filename (if the HTML value is a file path)
func (i HTML) Ext() string {
	return strings.ToLower(filepath.Ext(string(i)))
}

// Safe returns template.HTML to bypass HTML escaping (use carefully)
func (i HTML) Safe() template.HTML {
	return template.HTML(i)
}

// IsHTML returns true if the content looks like HTML markup
func (i HTML) IsHTML() bool {
	s := strings.TrimSpace(string(i))

	// Simple & fast detection
	if strings.HasPrefix(s, "<") && strings.HasSuffix(s, ">") {
		return true
	}

	// Or detect common tags
	if strings.Contains(s, "<") && strings.Contains(s, ">") {
		return true
	}

	return false
}
