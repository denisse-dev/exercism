// Package raindrops converts a number to a string, the contents of which depend
// on the number's factors.
package raindrops

import (
	"strconv"
)

var raindropSpeak = []struct {
	num    int
	speech string
}{
	{num: 3, speech: "Pling"},
	{num: 5, speech: "Plang"},
	{num: 7, speech: "Plong"},
}

// Convert transforms an integer to a string using randropSpeak rules.
func Convert(num int) (speech string) {
	for _, raindrop := range raindropSpeak {
		if num%raindrop.num == 0 {
			speech += raindrop.speech
		}
	}

	if speech == "" {
		speech = strconv.Itoa(num)
	}

	return speech
}
