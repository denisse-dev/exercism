package anagram

import (
	"reflect"
	"strings"
	"unicode"
)

// Alphabet stores letters and the number of incidences in a word.
type Alphabet map[rune]int

func wordCounter(s string) Alphabet {
	alphabet := make(Alphabet)
	for _, letter := range s {
		if unicode.IsLetter(letter) {
			alphabet[unicode.ToLower(letter)]++
		}
	}
	return alphabet
}

// Detect compares a list of words to a word to see if they're anagrams.
func Detect(word string, candidates []string) (anagrams []string) {
	wordMap := wordCounter(word)
	for _, candidate := range candidates {
		if len(word) != len(candidate) || strings.EqualFold(word, candidate) {
			continue
		}
		candidateMap := wordCounter(candidate)
		if reflect.DeepEqual(wordMap, candidateMap) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
