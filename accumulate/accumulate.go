// Package accumulate provides accumulate methods that the caller can use to accumulate.
package accumulate

type accumulateFunc func(string) string

// Accumulate takes a slice of strings s and an accumulateFunc fn,
// it'll then execute fn on each element in s and return the result.
func Accumulate(s []string, fn accumulateFunc) []string {
	result := make([]string, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}
