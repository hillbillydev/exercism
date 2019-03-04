// Package isbn provides functions for the ISBN specification.
package isbn

import (
	"errors"
	"strconv"
	"unicode"
)

var errInvalidISBN = errors.New("not a proper ISBN")

// Check validates the incoming input and determine if it's a valid ISBN.
func Check(s string) bool {
	is, err := toIntSlice(s)
	if err != nil {
		return false
	}

	return checksum(is)%11 == 0
}

func checksum(slice []int) int {
	var sum int
	for i, x := range slice {
		sum += x * (10 - i)
	}
	return sum
}

func toIntSlice(s string) ([]int, error) {
	var result []int
	for i, r := range s {
		// if we're at the last position we check if we have a check character.
		if len(s)-1 == i && r == 'X' {
			// X means it's a 10.
			result = append(result, 10)
			break
		}
		if unicode.IsLetter(r) {
			return nil, errInvalidISBN
		}

		number, err := strconv.Atoi(string(r))
		if err != nil {
			continue
		}
		result = append(result, number)
	}

	if len(result) != 10 {
		return nil, errInvalidISBN
	}

	return result, nil
}
