// Package proverb helps out with different proverbial rhymes.
package proverb

import "fmt"

// Proverb give it some rhymes and it will give you some badass rhymes.
func Proverb(rhymes []string) []string {
	var result []string
	for i, rhyme := range rhymes {
		if i == len(rhymes)-1 {
			result = append(result, fmt.Sprintf("And all for the want of a %s.", rhymes[0]))
			break
		}
		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme, rhymes[i+1]))
	}
	return result
}
