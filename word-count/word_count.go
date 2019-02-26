// Package wordcount counts words in an sentence.
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a map that puts each word as they key
// and the amount of times tha word happen has the value.
type Frequency map[string]int

// WordCount counts the words of an sentence.
func WordCount(sentence string) Frequency {
	result := make(Frequency)
	reg := regexp.MustCompile(`\w+('\w+)?`)

	for _, word := range reg.FindAllString(strings.ToLower(sentence), -1) {
		result[word]++
	}

	return result
}
