package types

import "strings"

type Email string

func (e Email) String() string  { return string(e) }
func (e Email) Kind() FieldType { return FieldEmail }
func (e Email) IsValid() bool {
	return strings.Contains(string(e), "@") && strings.Contains(string(e), ".")
}
