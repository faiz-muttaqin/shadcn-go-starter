package util

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatIndonesianPhone(phone string) string {
	// Remove non-digits
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(phone, "")

	// If starts with 62 or +62
	if strings.HasPrefix(digits, "62") {
		return fmt.Sprintf("+62 %s %s %s",
			digits[2:5], digits[5:9], digits[9:])
	}

	// Else, split 4-4-4 or whatever remains
	switch {
	case len(digits) > 8:
		return fmt.Sprintf("%s %s %s", digits[0:4], digits[4:8], digits[8:])
	case len(digits) > 4:
		return fmt.Sprintf("%s %s", digits[0:4], digits[4:])
	default:
		return digits
	}
}
