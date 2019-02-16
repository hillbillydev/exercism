// Package raindrops determine the sound from the weight of raindrops. SCIENCE!
package raindrops

import "strconv"

// Convert takes number i, it find all the prime factors it has.
// It then check if it contains 3,5 or 7.
// Each one determine the output of the function.
// - If the number has 3 as a factor, output 'Pling'.
// - If the number has 5 as a factor, output 'Plang'.
// - If the number has 7 as a factor, output 'Plong'.
// - If the number does not have 3, 5, or 7 as a factor,
func Convert(i int) string {
	var result string

	if i%3 == 0 {
		result += "Pling"
	}
	if i%5 == 0 {
		result += "Plang"
	}
	if i%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		return strconv.Itoa(i)
	}

	return result
}
