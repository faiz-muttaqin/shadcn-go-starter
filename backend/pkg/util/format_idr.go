package util

import (
	"fmt"
	"strings"
)

// FormatIDR formats an integer as Indonesian currency with thousands separators.
func FormatIDR(value int) string {
	isNegative := value < 0
	if isNegative {
		value = -value // Make the value positive for formatting
	}

	// Convert the integer to a string
	valueStr := fmt.Sprintf("%d", value)
	length := len(valueStr)

	// Insert points every three digits from the right
	var result strings.Builder
	for i, digit := range valueStr {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteString(".")
		}
		result.WriteRune(digit)
	}

	// Add a negative sign if the value was negative
	if isNegative {
		return "-Rp" + result.String()
	}

	return "Rp" + result.String()
}
