package atbash

import (
	"strings"
	"unicode"
)

// Atbash cipher is a simple substitution cipher that relies on
// transposing all the letters in the alphabet such that the resulting
// alphabet is backwards. The first letter is replaced with the last
// letter, the second with the second-last, and so on.
func Atbash(s string) string {
	var b strings.Builder

	s = strings.NewReplacer(
		" ", "",
		",", "",
		".", "",
	).Replace(strings.ToLower(s))

	for i, r := range s {
		// Every fifth element, but ignore the first one.
		if i%5 == 0 {
			b.WriteRune(' ')
		}

		if unicode.IsDigit(r) {
			// Should not find anything in the alphabeth,
			// instead I should only add it to the result.
			b.WriteRune(r)
			continue
		}

		b.WriteRune('z' + 'a' - r)
	}

	return strings.Trim(b.String(), " ")
}
