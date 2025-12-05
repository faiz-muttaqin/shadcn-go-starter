package util

import (
	"net/url"
	"path"
	"strings"
)

func GetPathOnly(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "/"
	}

	// Tambahkan scheme dummy jika tidak ada scheme
	if !strings.Contains(raw, "://") {
		raw = "http://" + raw
	}

	u, err := url.Parse(raw)
	if err != nil {
		// fallback: jika parse gagal, treat sebagai path biasa
		if !strings.HasPrefix(raw, "/") {
			return "/" + raw
		}
		return raw
	}

	// Ambil path dari URL
	p := u.Path
	if p == "" {
		p = "/"
	}

	// Pastikan normalized (/app bukan /app/)
	return path.Clean(p)
}
