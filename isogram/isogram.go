package isogram

import "unicode"

// IsIsogram tells if a string is an isogram.
func IsIsogram(word string) bool {
	visited := make(map[rune]bool)
	for _, chr := range word {
		if chr == ' ' || chr == '-' {
			continue
		}
		if visited[unicode.ToLower(chr)] == true {
			return false
		}
		visited[unicode.ToLower(chr)] = true
	}
	return true
}
