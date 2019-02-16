// Package etl extract transform load.
package etl

import (
	"strings"
	"unicode"
)

// Transform transform old data to newer fancier data.
func Transform(m map[int][]string) map[string]int {
	result := make(map[string]int)
	for _, s := range m {
		for _, v := range s {
			for _, s := range strings.ToLower(v) {
				result[string(s)] += value(s)
			}
		}
	}
	return result
}

func value(r rune) int {
	switch unicode.ToUpper(r) {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}
