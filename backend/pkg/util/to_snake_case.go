package util

import (
	"strings"
	"unicode"
)

// ToSnakeCase converts a camel case string to snake case
func ToSnakeCase(s string) string {
	var result strings.Builder

	for i, r := range s {
		// Check if the rune is uppercase
		if unicode.IsUpper(r) {
			// If it's the first rune, just append lowercase
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
