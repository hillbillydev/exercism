package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	r.name = generateRandomName()

	return r.name, nil
}

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

