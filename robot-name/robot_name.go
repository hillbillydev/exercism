// Package robotname handles the generation of robots.
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

var generator = newNameGenerator()

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
	name, err := generator.next()
	r.name = name

	return name, err
}

// Reset resets the name of the robot.
// Don't worry by calling for a name you'll get a new back.
func (r *Robot) Reset() {
	r.name = ""
}

type nameGenerator struct {
	names  []string
	cursor int
}

func (g *nameGenerator) next() (string, error) {
	if g.cursor == len(g.names) {
		return "", errors.New("error: names exhausted")
	}

	next := g.names[g.cursor]
	g.cursor++

	return next, nil
}

func newNameGenerator() *nameGenerator {
	return &nameGenerator{
		names: generateNames(),
	}
}

func generateNames() []string {
	var (
		result []string

		first  = 'A'
		second = 'A'
		number = 0
	)

	for first <= 'Z' {
		result = append(result, fmt.Sprintf("%s%s%03d", string(first), string(second), number))

		if number != 999 {
			number++
			continue
		}
		number = 0
		second++

		if second <= 'Z' {
			continue
		}

		first++
		second = 'A'
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result
}
