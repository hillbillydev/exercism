// Package isogram helps you out figuring out if a string is a isogram.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if the passed in string is an Isogram.
// it determines that by looking at each rune in the string
// if the rune count goes over 1 we know it's not a Isogram.
func IsIsogram(s string) bool {
	runes := make(map[rune]bool)
	for _, r := range strings.ToLower(s) {
		if !unicode.IsLetter(r) {
			continue
		}
		if v := runes[r]; v {
			return false
		}

		runes[r] = true
	}

	return true
}

// 200000	      6786 ns/op	    1095 B/op	      11 allocs/op
