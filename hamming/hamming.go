// Package hamming calculates the Hamming difference bewtween to DNA strands
package hamming

import (
	"errors"
)

// Distance returns the number of differences between two homologous DNA strans
// taken from different genomes with a common ancestor.
func Distance(a, b string) (mutation int, err error) {
	if len(a) != len(b) {
		return -1, errors.New("the strings are not of equal length")
	}

	for i := range a {
		if a[i] != b[i] {
			mutation++
		}
	}
	return mutation, nil
}
