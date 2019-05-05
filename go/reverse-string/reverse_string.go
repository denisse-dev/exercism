package reverse

import "strings"

// String returns a string in inverse order
func String(word string) string {
	var reversed strings.Builder
	letters := []rune(word)
	last := len(letters) - 1

	for i := 0; i < len(letters); i++ {
		reversed.WriteRune(letters[last-i])
	}

	return reversed.String()
}
