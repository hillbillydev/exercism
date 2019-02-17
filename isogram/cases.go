// Package isogram helps you out figuring out if a string is a isogram.
package isogram

import "unicode"

// IsIsogram determines if the passed in string is an Isogram.
// it determines that by looking at each rune in the string
// if the rune count goes over 1 we know it's not a Isogram.
func IsIsogram(s string) bool {
	runeMap := make(map[rune]struct{})
	for _, r := range s {
		r = unicode.ToLower(r)
		if !unicode.IsLetter(r) {
			continue
		}
		if _, exists := runeMap[r]; exists {
			return false
		}

		runeMap[r] = struct{}{}
	}

	return true
}
