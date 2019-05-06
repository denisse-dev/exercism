package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is a type used to store a word and it's occurrences.
type Frequency map[string]int

// WordCounts counts the occurrences of each word in a given phrase.
func WordCount(phrase string) Frequency {
	freq := make(Frequency)
	phrase = strings.ToLower(phrase)
	condition := func(c rune) bool {
		return c != '\'' && !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	for _, word := range strings.FieldsFunc(phrase, condition) {
		freq[strings.Trim(word, "'")]++
	}

	return freq
}
