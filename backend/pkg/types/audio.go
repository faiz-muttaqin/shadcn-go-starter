package types

import (
	"path/filepath"
	"strings"
)

type Audio string

func (a Audio) String() string  { return string(a) }
func (a Audio) Kind() FieldType { return FieldAudio }

func (a Audio) Ext() string {
	return strings.ToLower(filepath.Ext(string(a)))
}

func (a Audio) IsAudio() bool {
	switch a.Ext() {
	case ".mp3", ".wav", ".aac", ".flac", ".ogg", ".m4a", ".amr":
		return true
	}
	return false
}
