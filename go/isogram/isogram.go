package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(word string) (isogram bool) {
	word = strings.ToLower(word)

	for i, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}
		if strings.ContainsRune(word[i+1:], letter) {
			return false
		}
	}

	return true
}
