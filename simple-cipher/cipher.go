package cipher

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Caesar struct {
	shifter []int
}

func NewCaesar() *Caesar {
	return &Caesar{
		shifter: []int{3},
	}
}

func NewShift(i int) *Caesar {
	if i == 0 || i >= 26 || i <= -26 {
		return nil
	}

	return &Caesar{
		shifter: []int{i},
	}
}

func (c *Caesar) getShift(i int) int {
	return c.shifter[i%len(c.shifter)]
}

func NewVigenere(key string) *Caesar {
	shifters, err := constructShifters(key)
	if err != nil {
		return nil
	}

	return &Caesar{shifter: shifters}
}

func (c *Caesar) Encode(s string) string {
	var result []rune

	for _, r := range strings.ToLower(s) {
		if !unicode.IsLetter(r) {
			continue
		}
		shift := c.getShift(len(result))

		result = append(result, wrappedRune(r+rune(shift)))
	}

	return string(result)
}

func (c *Caesar) Decode(s string) string {
	var result []rune
	for i, r := range s {
		shift := c.getShift(i)
		result = append(result, wrappedRune(r-rune(shift)))
	}

	return string(result)
}

func wrappedRune(r rune) rune {
	if r < 'a' {
		r = 'z' - 'a'%r + 1
	}

	if r > 'z' {
		r = 'a' - 1 + r%'z'
	}
	return r

}

func constructShifters(key string) ([]int, error) {
	var result []int
	if len(key) <= 2 {
		return nil, fmt.Errorf("Length of key can't be less then 3 but has a lenght of %d", len(key))
	}

	for _, r := range key {
		if !unicode.IsLetter(r) || unicode.IsUpper(r) {
			return nil, errors.New("error: the key can't non letter or uppercase")
		}
		shifterValue := int(r) - 'a'

		result = append(result, shifterValue)
	}

	return result, nil
}
