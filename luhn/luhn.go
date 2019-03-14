// Package luhn is a package that determine the validity of a Luhn sequence.
package luhn

import (
	"regexp"
	"strconv"
)

var trimSpaces = regexp.MustCompile(`[^ ]`).FindAllString

// Valid determines the validity of the input, if it is an luhn sequence we give back true otherwise false.
func Valid(input string) bool {
	var (
		sum     int
		trimmed = trimSpaces(input, -1)
		l       = len(trimmed)
	)
	if l <= 1 {
		return false
	}

	for i := l; i > 0; i-- {
		n, err := strconv.Atoi(trimmed[i-1])
		if err != nil {
			return false
		}

		if (l%2 == 0) == (i%2 == 0) {
			sum += n
			continue
		}

		n *= 2
		if n > 9 {
			n -= 9
		}
		sum += n
	}

	return sum%10 == 0
}
