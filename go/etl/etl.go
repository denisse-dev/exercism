package etl

import "unicode"

// Transform converts a legacy scrabble score to a shiny new scrabble score
func Transform(legacy map[int][]string) map[string]int {
	shiny := make(map[string]int)
	for key, values := range legacy {
		for _, value := range values {
			letter := []rune(value)
			shiny[string(unicode.ToLower(letter[0]))] = key
		}
	}

	return shiny
}
