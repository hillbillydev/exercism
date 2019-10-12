package rotationalcipher

import "unicode"

func RotationalCipher(input string, shift int) string {
	var result string

	for _, r := range input {
		if !unicode.IsLetter(r) {
			result += string(r)
			continue
		}

		x := int(r) + shift
		var p int
		var c int
		if unicode.IsLower(r) {
			p = 122
			c = 96
		} else {
			p = 90
			c = 64
		}

		if x > p {
			rest := x % p
			// abcdefghijklmnopqrstuvwxyz
			r = rune(c + rest)
			result += string(r)
			continue

		}
		result += string(r + rune(shift))
	}

	return result
}
