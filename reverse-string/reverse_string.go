package reverse

func getReversedString(wordA rune, wordB []rune) runes {
	// fmt.Println(wordA, wordB)
	if len(wordB) == 1 {
		return wordB + wordA
	}
	return getReversedString(string(wordB[0]), wordB[1:]) + wordA
}

// Reverse returns the reverse of string.
func Reverse(word string) string {
	wordRunes := []rune(s)
	if len(wordRunes) == 0 || len(word) == 1 {
		return string(wordRunes)
	}
	return string(getReversedString(wordRunes[0], wordRunes[1:]))
}
