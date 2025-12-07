package types

import (
	"html/template"
	"path/filepath"
	"strings"
)

type CSS template.CSS

func (i CSS) String() string  { return string(i) }
func (i CSS) Kind() FieldType { return FieldCSS }

// Ext returns the extension of a file path
func (i CSS) Ext() string {
	return strings.ToLower(filepath.Ext(string(i)))
}

// Safe returns template.CSS to bypass escaping
func (i CSS) Safe() template.CSS {
	return template.CSS(i)
}

// IsCSS tries to detect if content looks like CSS
func (i CSS) IsCSS() bool {
	s := strings.TrimSpace(string(i))

	// File extension detection
	if i.Ext() == ".css" {
		return true
	}

	// Detect CSS selectors or blocks
	if strings.Contains(s, "{") && strings.Contains(s, "}") {
		return true
	}

	// Detect common CSS syntax
	cssIndicators := []string{
		":", // property:value
		";", // rule ending
		"color",
		"margin",
		"padding",
		"display",
		"font",
		"background",
	}

	for _, key := range cssIndicators {
		if strings.Contains(s, key) {
			return true
		}
	}

	return false
}
