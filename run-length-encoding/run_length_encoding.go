package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	var (
		builder = strings.Builder{}
		count   int
	)

	for i, r := range input {
		next := '0'
		count++

		if i != len(input)-1 {
			next = rune(input[i+1])
		}

		if r == next {
			continue
		}
		if count > 1 {
			builder.WriteString(strconv.Itoa(count))
		}
		count = 0
		builder.WriteRune(r)
	}

	return builder.String()
}

func RunLengthDecode(input string) string {
	var (
		result = strings.Builder{}
		digits = ""
	)

	for _, r := range input {
		if unicode.IsDigit(r) {
			digits += string(r)
			continue
		}

		digit, err := strconv.Atoi(digits)
		if err != nil {
			result.WriteRune(r)
			continue
		}

		result.WriteString(strings.Repeat(string(r), digit))

		digits = ""
	}

	return result.String()
}
