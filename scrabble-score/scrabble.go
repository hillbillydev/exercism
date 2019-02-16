// Package scrabble let you play and have fun with your friend in the game of scrabble!
package scrabble

import "unicode"

// Score determines the score you receive when you finnish a word.
// Ex. "cabbage" -> 14 points.
func Score(s string) int {
	var result int
	for _, r := range s {
		result += value(r)
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
