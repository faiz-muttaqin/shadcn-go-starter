package types

import (
	"path/filepath"
	"strings"
)

type Image string

func (i Image) String() string  { return string(i) }
func (i Image) Kind() FieldType { return FieldImage }

func (i Image) Ext() string {
	return strings.ToLower(filepath.Ext(string(i)))
}

func (i Image) IsImage() bool {
	switch i.Ext() {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".bmp", ".svg":
		return true
	}
	return false
}
