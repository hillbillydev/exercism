package bob

import (
	"strings"
	"unicode"
)

// Hey represent weird bob, a guy who is quite boring to talk to.
func Hey(remark string) string {
	if isCapitalSentence(remark) && strings.HasSuffix(remark, "?") {
		return "Calm down, I know what I'm doing!"
	} else if strings.HasSuffix(strings.TrimSpace(remark), "?") {
		return "Sure."
	} else if isCapitalSentence(remark) {
		return "Whoa, chill out!"
	} else if strings.TrimSpace(remark) == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

func isCapitalSentence(s string) bool {
	noLetter := true
	for _, r := range s {
		if unicode.IsLetter(r) {
			noLetter = false
		}
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			return false
		}
	}

	if noLetter {
		return false
	}

	return true
}
