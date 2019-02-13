// Package acronym helps out with acronyms.
package acronym

import (
	"strings"
)

// Abbreviate takes a string s and figures out what the acronym is.
// "Ruby on Rails" -> "ROR"
func Abbreviate(s string) string {
	builder := strings.Builder{}
	s = strings.NewReplacer("'", "", "-", " ").Replace(s)
	for _, word := range strings.Split(s, " ") {
		if len(word) == 0 {
			continue
		}
		builder.WriteString(word[:1])
	}
	return strings.ToUpper(builder.String())
}

// NEW 100000	     15599 ns/op	   53712 B/op	      61 allocs/op
