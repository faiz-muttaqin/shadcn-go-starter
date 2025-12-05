package util

import (
	"regexp"
	"unicode"
)

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
func IsNumericRegex(s string) bool {
	matched, _ := regexp.MatchString(`^\d+$`, s)
	return matched
}
