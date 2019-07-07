package cryptosquare

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func Encode(input string) string {
	var normalized string
	for _, r := range strings.ToLower(input) {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		}
		normalized += string(r)
	}

	col := int(math.Ceil(math.Sqrt(float64(len(normalized)))))

	chunks := make([]string, col)
	for i := 0; i < col; i++ {
		for j := i; j < len(normalized); j += col {
			x := string(normalized[j])
			chunks[i] += x
		}
	}

	var max int
	for i, c := range chunks {
		if len(c) >= max {
			max = len(c)
			continue
		}
		chunks[i] = fmt.Sprintf("%-*s", max, c)
	}

	return strings.Join(chunks, " ")
}
