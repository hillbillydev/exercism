// Package proverb helps out with different proverbial rhymes.
package proverb

import "fmt"

// Proverb give it some rhymes and it will give you some badass rhymes.
func Proverb(rhymes []string) []string {
	result := make([]string, len(rhymes))
	for i, rhyme := range rhymes {
		if i == len(rhymes)-1 {
			result[i] = fmt.Sprintf("And all for the want of a %s.", rhymes[0])
			break
		}
		result[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme, rhymes[i+1])
	}
	return result
}
