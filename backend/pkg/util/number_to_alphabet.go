package util

func NumberToAlphabet(n int) string {
	result := ""
	for n > 0 {
		// Calculate remainder when n is divided by 26
		remainder := (n - 1) % 26
		// Convert remainder to a letter ('A' to 'Z')
		char := 'A' + rune(remainder)
		// Prepend the character to the result
		result = string(char) + result
		// Update n to the next column (reduce it)
		n = (n - 1) / 26
	}
	return result
}
