// Package pangram gives good functionality when working with pangrams.
package pangram

import (
	"strings"
	"unicode"
)

// IsPangram determines if the sentence contains the whole albabeth.
func IsPangram(sentence string) bool {
	runes := make(map[rune]int)
	for _, r := range strings.ToLower(sentence) {
		if !unicode.IsLetter(r) {
			continue
		}
		runes[r]++
	}
	return len(runes) == 26
}
