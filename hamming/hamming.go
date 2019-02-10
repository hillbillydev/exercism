// Package hamming provides functions to work with the difference between DNA strands.
package hamming

import "errors"

// ErrHammingDistanceNotEqual happens when DNA stands are not equal.
var ErrHammingDistanceNotEqual = errors.New("length between both dna's most be equal to each other")

// Distance calculates the hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, ErrHammingDistanceNotEqual
	}
	var count int
	sa := []rune(a)
	sb := []rune(b)

	for i := 0; i < len(a); i++ {
		if sa[i] != sb[i] {
			count++
		}
	}
	return count, nil
}
