package util

import (
	"os"
	"path/filepath"
	"strings"
)

func IsDevMode() bool {
	exe, _ := os.Executable()
	exe = filepath.ToSlash(exe)
	if strings.HasSuffix(filepath.Dir(exe), "/tmp") || strings.HasSuffix(filepath.Dir(exe), "/.tmp") {
		return true
	}

	// air → build di tmp_dir lokal
	if strings.Contains(exe, "/tmp/") || strings.Contains(strings.ToLower(exe), "/temp/") {
		return true
	}
	if strings.Contains(exe, "\\tmp\\") || strings.Contains(strings.ToLower(exe), "\\temp\\") {
		return true
	}
	if strings.Contains(exe, "/go-build") || strings.Contains(exe, "\\go-build") {
		return true
	}

	return false
}
