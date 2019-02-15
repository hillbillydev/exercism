// Package dna works and operates on dna's.
package dna

import (
	"errors"
)

// ErrInvalidNucleotide is an error that occurs when the caller have an invalid nucleotide in the strand.
var ErrInvalidNucleotide = errors.New("invalid nucleotide")

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA struct {
	strand string
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := newHistogram()
	for _, rune := range d.strand {
		if ok := validateNucleotide(rune); !ok {
			return Histogram{}, ErrInvalidNucleotide
		}
		h[rune]++
	}
	return h, nil
}

func validateNucleotide(r rune) bool {
	switch r {
	case 'A', 'C', 'G', 'T':
		return true
	default:
		return false
	}
}

func newHistogram() Histogram {
	return Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
}
