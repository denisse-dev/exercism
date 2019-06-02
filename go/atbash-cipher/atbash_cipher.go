package atbash

import (
	"strings"
	"unicode"
)

// Atbash returns an encrypted string using the ancient Atbash Cipher.
func Atbash(decrypted string) string {
	var encrypted strings.Builder
	var count int

	for _, letter := range decrypted {
		if unicode.IsSpace(letter) || unicode.IsPunct(letter) {
			continue
		}
		if count == 5 {
			encrypted.WriteRune(' ')
			count = 0
		}
		if unicode.IsDigit(letter) {
			encrypted.WriteRune(letter)
			count++
			continue
		}
		diff := 'z' - unicode.ToLower(letter)
		encrypted.WriteRune('a' + diff)
		count++
	}

	return encrypted.String()
}
