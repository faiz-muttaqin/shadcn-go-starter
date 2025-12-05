package util

import "fmt"

func FormatFileSize(size int64) string {
	switch {
	case size > 1<<30:
		return fmt.Sprintf("%.2f GB", float64(size)/(1<<30))
	case size > 1<<20:
		return fmt.Sprintf("%.2f MB", float64(size)/(1<<20))
	case size > 1<<10:
		return fmt.Sprintf("%.2f KB", float64(size)/(1<<10))
	default:
		return fmt.Sprintf("%d B", size)
	}
}
