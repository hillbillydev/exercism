// Package acronym helps out with acronyms.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate takes a string s and figures out what the acronym is.
// "Ruby on Rails" -> "ROR"
func Abbreviate(s string) string {
	builder := strings.Builder{}
	// Not sure about the replace here, it solves the issue, but it seems like a hack?
	// The problem is that "Halley's Comet" become HSC.
	normalized := strings.Replace(strings.ToLower(s), "'", "", -1)
	for _, rune := range strings.Title(normalized) {
		if unicode.IsLetter(rune) && unicode.IsUpper(rune) {
			builder.WriteRune(unicode.ToUpper(rune))
		}
	}
	return builder.String()
}
