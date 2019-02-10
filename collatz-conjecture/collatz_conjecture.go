//Package collatzconjecture helps out with the Collatz Conjecture problem.
package collatzconjecture

import (
	"errors"
)

// ErrInvalidArgument is an error that happens when the caller gives an invalid argument.
var ErrInvalidArgument = errors.New("argument can't be zero or negative")

// CollatzConjecture determine how many steps it take to solve the CollatzConjecture.
func CollatzConjecture(i int) (int, error) {
	var steps int
	if i <= 0 {
		return 0, ErrInvalidArgument
	}

	for i != 1 {
		steps++
		if i%2 == 0 {
			i = i / 2
			continue
		}

		i = i*3 + 1
	}

	return steps, nil
}
