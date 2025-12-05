package types

import (
	"path/filepath"
	"strings"
)

type Archive string

func (a Archive) String() string  { return string(a) }
func (a Archive) Kind() FieldType { return FieldArchive }

func (a Archive) Ext() string {
	return strings.ToLower(filepath.Ext(string(a)))
}

func (a Archive) IsArchive() bool {
	switch a.Ext() {
	case ".zip", ".rar", ".7z", ".tar", ".gz":
		return true
	}
	return false
}
