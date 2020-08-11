// Package robotname handles the generation of robot names.
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

var used = map[string]bool{}

// Robot, makes sure humans does not need to work anymore.
type Robot struct {
	name string
}

// Name return the next name.
func (r *Robot) Name() (string, error) {
	if len(used) == 26*26*10*10*10 {
		return "", errors.New("error: names exhausted")
	}

	if r.name != "" {
		return r.name, nil
	}

	r.name = generateName()
	for used[r.name] {
		r.name = generateName()
	}
    used[r.name] = true

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
    r1 := random.Intn(26) + 'A'
	r2 := random.Intn(26) + 'A'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}

