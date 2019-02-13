package raindrops

import (
	"strconv"
	"strings"
)

// Convert takes number i, it find all the prime factors it has.
// It then check if it contains 3,5 or 7.
// Each one determine the output of the function.
// - If the number has 3 as a factor, output 'Pling'.
// - If the number has 5 as a factor, output 'Plang'.
// - If the number has 7 as a factor, output 'Plong'.
// - If the number does not have 3, 5, or 7 as a factor,
func Convert(i int) string {
	factors := primeFactors(i)
	builder := strings.Builder{}

	if containsFactor(factors, 3) {
		builder.WriteString("Pling")
	}

	if containsFactor(factors, 5) {
		builder.WriteString("Plang")
	}

	if containsFactor(factors, 7) {
		builder.WriteString("Plong")
	}

	if builder.Len() == 0 {
		return strconv.Itoa(i)
	}
	return builder.String()
}

func primeFactors(n int) []int {
	var factors []int
	// Get the number of 2s that divide n
	for n%2 == 0 {
		factors = append(factors, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

// checks if we can find a certain factor within the slice.
func containsFactor(s []int, factor int) bool {
	for _, number := range s {
		if number == factor {
			return true
		}
	}
	return false
}
