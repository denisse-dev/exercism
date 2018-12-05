// Package dna computes how many times each nucleotde occurs in the DNA string.
package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (h Histogram, err error) {
	var G, C, T, A int

	for i := range d {
		switch d[i] {
		case 'G':
			G++
		case 'C':
			C++
		case 'T':
			T++
		case 'A':
			A++
		default:
			return nil, errors.New("invald nucleotid")
		}
	}

	return Histogram{'G': G, 'C': C, 'T': T, 'A': A}, nil
}
