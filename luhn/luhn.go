// Package luhn is a package that determine the validity of a Luhn sequence.
package luhn

import (
	"strconv"
	"unicode"
)

// Valid determines the validity of the input, if it is an luhn sequence we give back true otherwise false.
func Valid(input string) bool {
	luhns, ok := toValidLuhns(input)
	if !ok {
		return false
	}

	var sum int
	for i := len(luhns); i > 0; i-- {
		current := luhns[i-1]
		if (len(luhns)%2 == 0) == (i%2 == 0) {
			sum += current
			continue
		}

		current *= 2
		if current > 9 {
			current -= 9
		}

		sum += current
	}

	return sum%10 == 0
}

func toValidLuhns(s string) ([]int, bool) {
	var luhns []int
	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}

		nr, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, false
		}

		luhns = append(luhns, nr)
	}

	if len(luhns) <= 1 {
		return nil, false
	}

	return luhns, true
}
