package util

import "unicode"

func AddSpaceBeforeUppercase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) && i != 0 {
			result = append(result, ' ')
		}
		result = append(result, r)
	}
	return string(result)
}
