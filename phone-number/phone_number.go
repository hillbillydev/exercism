package phonenumber

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Number takes multiple scrambled formats of an phone number
// it then parses each rune and tries to figure out if it's a valid phone number.
// Example (223) 456-7890 -> 2234567890
func Number(number string) (string, error) {
	result := strings.Builder{}

	for _, r := range number {
		if unicode.IsLetter(r) {
			return "", fmt.Errorf("NoLettersErr: didn't expect any errors but got %q", string(r))
		}
		if !unicode.IsDigit(r) {
			continue
		}

		digit, _ := strconv.Atoi(string(r))
		if result.Len() == 0 || result.Len() == 3 {
			if digit < 2 {
				continue
			}
		}

		result.WriteRune(r)
	}

	if result.Len() != 10 {
		return "", fmt.Errorf("InvalidNumberErr: expected an length of %d but got %d", 10, result.Len())
	}

	return result.String(), nil
}

// AreaCode take a number and extracts the AreaCode from it.
// Example 2234567890 -> 223
func AreaCode(number string) (string, error) {
	number, err := Number(number)
	if err != nil {
		return "", err
	}

	return number[0:3], nil
}

// Format takes a number and formats it.
// Example 2234567890 -> (223) 456-7890.
func Format(number string) (string, error) {
	number, err := Number(number)
	if err != nil {
		return "", err
	}
	// (613) 995-0253
	return fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:10]), nil
}
