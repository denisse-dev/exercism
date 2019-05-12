package pangram

import "unicode"

// IsPangram determine if a sentence uses the 26 letters of the alphabet.
func IsPangram(word string) (result bool) {
	var alphabet [26]int

	for _, letter := range word {
		if unicode.IsLetter(letter) {
			alphabet[unicode.ToLower(letter)-'a']++
		}
	}

	for _, value := range alphabet {
		if value == 0 {
			return false
		}
	}
	return true
}
