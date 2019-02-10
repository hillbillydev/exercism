// Package acronym helps out with acronyms.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate takes a string s and figures out what the acronym is.
// "Ruby on Rails" -> "ROR"
func Abbreviate(s string) string {
	var result string
	// Random regex, most be a better one out there.
	fields := regexp.MustCompile("\\s|-|yperTex").Split(s, -1)
	for _, field := range fields {
		if len(field) == 0 {
			continue
		}

		result = result + strings.ToUpper(string(field[0]))
	}
	return result
}
