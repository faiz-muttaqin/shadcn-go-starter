package util

import (
	"path/filepath"
)

// call by input runtime.Caller(0)
func ThisFileDir(pc uintptr, file string, line int, ok bool) string {
	// _, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}
