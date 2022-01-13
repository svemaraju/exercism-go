package reverse

// getReversedString is recursive function
// to reverse a string
func getReversedString(wordA rune, wordB []rune) string {
    if len(wordB) == 0 {
        return string(wordA)
    }
	if len(wordB) == 1 {
		return string(wordB[0]) + string(wordA)
	}
	return getReversedString(wordB[0], wordB[1:]) + string(wordA)
}

// Reverse returns the reverse of string.
func Reverse(word string) string {
	wordRunes := []rune(word)
	if len(wordRunes) == 0 || len(word) == 1 {
		return string(wordRunes)
	}
	return string(getReversedString(wordRunes[0], wordRunes[1:]))
}
