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
	numbers := 100 + rand.Intn(999-100)

	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 2)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return fmt.Sprintf("%s%d", string(b), numbers)
}
