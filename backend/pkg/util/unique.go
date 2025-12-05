package util

func Unique[T comparable](input []T) []T {
	seen := make(map[T]bool)
	var result []T

	for _, v := range input {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}
