package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher is an implementation of the ancient rotational cipher.
func RotationalCipher(input string, rot int) string {
	if rot%26 == 0 {
		return input
	}

	var ciphertext strings.Builder
	var first, last rune
	for _, letter := range input {
		if !unicode.IsLetter(letter) {
			ciphertext.WriteRune(letter)
			continue
		}

		if unicode.IsUpper(letter) {
			first = 'A'
			last = 'Z'
		} else {
			first = 'a'
			last = 'z'
		}

		diff := last - letter
		if rune(rot) > diff {
			letter = first - diff - 1
		}

		ciphertext.WriteRune(letter + rune(rot))
	}

	return ciphertext.String()
}
