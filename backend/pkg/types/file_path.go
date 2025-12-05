package types

import (
	"path/filepath"
	"strings"
)

type File string

func (f File) String() string  { return string(f) }
func (f File) Kind() FieldType { return FieldFile }

func (f File) Ext() string {
	return strings.ToLower(filepath.Ext(string(f)))
}
