// Package robotname handles the generation of robots.
package robotname

import (
	"fmt"
	"math/rand"
)

// Robot represents the core struct in this package.
type Robot struct {
	name string
}

// Name gives back the name of the robot,
// if it doesn't have a name it generates one.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	r.name = generateRandomName()

	return r.name, nil
}

// Reset removes the old name and generate a new one.
func (r *Robot) Reset() {
	r.name = generateRandomName()
}

func generateRandomName() string {
	var (
		numbers = 100 + rand.Intn(999-100)
		letters = []rune{randomRune(), randomRune()}
	)

	return fmt.Sprintf("%s%d", string(letters), numbers)
}

func randomRune() rune {
	return 65 + rune(rand.Intn(91-65))
}
