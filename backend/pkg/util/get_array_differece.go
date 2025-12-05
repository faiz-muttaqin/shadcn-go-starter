package util

func GetArrayDifference(strings1, strings2 []string) []string {
	// Create a map for strings2 for fast lookup
	lookup := make(map[string]bool)
	for _, s := range strings2 {
		lookup[s] = true
	}

	// Now collect elements in strings1 that are not in strings2
	var result []string
	for _, s := range strings1 {
		if !lookup[s] {
			result = append(result, s)
		}
	}

	return result
}
