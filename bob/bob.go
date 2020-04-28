// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings";
	"unicode"
)

func IsUpper(remark string) bool {
	numLetters := 0
	for _, chr := range remark {
		if !unicode.IsUpper(chr) && unicode.IsLetter(chr) {
			return false
		}
		if unicode.IsLetter(chr) {
			numLetters += 1
		}
	}
	if numLetters == 0 { 
		return false
	}
	return true
}
// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	qmark := "?"
	remark = strings.TrimSpace(remark)
	if len(remark) == 0{
		return "Fine. Be that way!"
	} else if IsUpper(remark) && remark[len(remark)-1:] == qmark {
		return "Calm down, I know what I'm doing!"
	} else if remark[len(remark)-1:] == qmark {
		return "Sure."
	} else if IsUpper(remark) {
		return "Whoa, chill out!"
	} 
	return "Whatever."
}
