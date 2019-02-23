// Package anagram contains functions to determine if stuff is a anagram.
package anagram

import (
	"strings"
	"unicode"
)

// Detect takes a subject and candidates, it will then figure
// if you can scramble the subject to become one of those candidates
// ex, subject -> pam, candidates have a value "map" that would make that a match.
func Detect(subject string, candidates []string) []string {
	var (
		result       []string
		subjectRunes = runes(subject)
	)

	for _, v := range candidates {
		if !validateInputs(subject, v) {
			continue
		}

		if !compareSubjectAndCandidate(subjectRunes, runes(v)) {
			continue
		}

		result = append(result, v)
	}
	return result
}

func validateInputs(s, compare string) bool {
	if len(s) != len(compare) {
		// It's not an anagram, bail.
		return false
	}

	if strings.ToLower(s) == strings.ToLower(compare) {
		// same words can't be an anagram.
		return false
	}
	return true
}

func runes(s string) map[rune]int {
	result := make(map[rune]int)
	for _, r := range s {
		result[unicode.ToLower(r)]++
	}
	return result
}

func compareSubjectAndCandidate(subject, candidate map[rune]int) bool {
	for key, count := range candidate {
		if subject[key] != count {
			return false
		}
	}
	return true
}
