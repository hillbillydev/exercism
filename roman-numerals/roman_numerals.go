// Package romannumerals transforms number to the old school way of
// representing numbers.
package romannumerals

import "errors"

// Roman a struct that has a number and the visual number as a roman number.
type Roman struct {
	number int
	arabic string
}

// Convert takes a normal int and converts it to an roman numeral.
func Convert(i int) (string, error) {
	if i <= 0 || i > 3000 {
		return "", errors.New("invalid argument")
	}
	var romans = []Roman{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result string
	for _, r := range romans {
		for i >= r.number {
			result += r.arabic
			i -= r.number
		}
	}

	return result, nil
}
