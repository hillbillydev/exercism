package triangle

import (
	"math"
	"sort"
)

// Kind represent different kinds of triangles
type Kind int32

// Different kinds of triangles
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}

	s := []float64{a, b, c}
	sort.Float64s(s)

	if s[0] == s[2] {
		return Equ
	} else if s[0] == s[1] || s[1] == s[2] {
		return Iso
	} else {
		return Sca
	}
}

func isTriangle(a, b, c float64) bool {
	if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return false
	}

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return false
	}

	if a <= 0 && b <= 0 && c <= 0 {
		return false
	}
	if (a+b >= c) && (a+c >= b) && (b+c >= a) {
		return true
	}

	return false
}

// func countSameSides(a, b, c float64) bool {

// 	return count
// }
