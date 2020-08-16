package darts

import (
	"math"
)

// Score takes x,y of the throwned dart
// and gives back the total score of that throw.
func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)

	if distance <= 1 {
		return 10
	}
	if distance <= 5 {
		return 5
	}
	if distance <= 10 {
		return 1
	}

	return 0
}
