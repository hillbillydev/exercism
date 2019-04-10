// Package grains helps out counting grains on a chess board.
package grains

import (
	"errors"
)

// Total gives the all the grains on each square on a chess board.
func Total() uint64 {
	sum := 0 << 63
	return uint64(^sum)
}

// Square take the number of the square and return the amount of grains on that square.
func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, errors.New("invalid argument")
	}
	// grains := math.Pow(2, float64(i-1))
	grains := 1 << uint(i-1)
	return uint64(grains), nil
}
