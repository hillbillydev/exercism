// Package pangram gives good functionality when working with pangrams.
package pangram

import (
	"strings"
)

// IsPangram checks that input is a pangram.
func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for r := 'a'; r <= 'z'; r++ {
		if !strings.ContainsRune(input, r) {
			return false
		}
	}
	return true
}
