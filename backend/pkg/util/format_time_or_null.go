package util

import "time"

func FormatTimeOrNull(t time.Time, layout string) string {
	if t.IsZero() {
		return "null"
	}
	return t.Format(layout)
}
