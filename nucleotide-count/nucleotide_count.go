// Package dna works and operates on dna's.
package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA represents a structure of the DNA.
type DNA struct {
	strand string
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, r := range d.strand {
		if _, exist := h[r]; !exist {
			return Histogram{}, fmt.Errorf("%s is not a valid nucleotide", string(r))
		}
		h[r]++
	}
	return h, nil
}
