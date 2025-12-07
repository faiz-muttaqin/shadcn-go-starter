package types

import (
	"html/template"
	"path/filepath"
	"strings"
)

type JS template.JS

func (i JS) String() string  { return string(i) }
func (i JS) Kind() FieldType { return FieldJS }

// Ext returns the extension of the filename (if the JS value is a file path)
func (i JS) Ext() string {
	return strings.ToLower(filepath.Ext(string(i)))
}

// Safe returns template.JS to bypass JS escaping (use carefully)
func (i JS) Safe() template.JS {
	return template.JS(i)
}

// IsJS tries to detect if content looks like JavaScript code
func (i JS) IsJS() bool {
	s := strings.TrimSpace(string(i))

	// File extension detection
	if i.Ext() == ".js" {
		return true
	}

	// Detect common JS syntax
	jsIndicators := []string{
		"function ",
		"()=>",
		"var ",
		"let ",
		"const ",
		"document.",
		"window.",
		"console.",
		"import ",
		"export ",
	}

	for _, key := range jsIndicators {
		if strings.Contains(s, key) {
			return true
		}
	}

	return false
}
