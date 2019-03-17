// Package luhn is a package that determine the validity of a Luhn sequence.
package luhn

import (
	"strconv"
	"strings"
)

// Valid determines the validity of the input, if it is an luhn sequence we give back true otherwise false.
func Valid(input string) bool {
	var (
		sum     int
		trimmed = strings.ReplaceAll(input, " ", "")
		parity  = len(trimmed) % 2
	)

	if len(trimmed) <= 1 {
		return false
	}

	for i := len(trimmed) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(trimmed[i]))
		if err != nil {
			return false
		}

		if i%2 == parity {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}

	return sum%10 == 0
}
