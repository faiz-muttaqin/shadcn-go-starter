package types

import (
	"strings"
	"unicode"
)

type Phone string

func (p Phone) String() string  { return string(p) }
func (p Phone) Kind() FieldType { return FieldPhone }

func (p Phone) Normalize() Phone {
	s := strings.TrimSpace(string(p))
	return Phone(strings.ReplaceAll(s, "-", ""))
}

func (p Phone) IsNumeric() bool {
	for _, r := range string(p) {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
