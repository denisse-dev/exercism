// Package proverb generates a relevant proverb similar to "For Want of a Nail"
package proverb

import (
	"fmt"
)

const (
	baseString = "For want of a %s the %s was lost."
	lastString = "And all for the want of a %s."
)

// Proverb given a slice of strings the correct proverb is generated
func Proverb(rhyme []string) (output []string) {
	for i := range rhyme {
		if i == len(rhyme)-1 || len(rhyme) == 1 {
			output = append(output, fmt.Sprintf(lastString, rhyme[0]))
		} else {
			output = append(output, fmt.Sprintf(baseString, rhyme[i], rhyme[i+1]))
		}
	}
	return
}
