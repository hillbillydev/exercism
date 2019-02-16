// Package etl extract transform load.
package etl

import "strings"

// Transform transform old data to newer fancier data.
func Transform(m map[int][]string) map[string]int {
	result := make(map[string]int)
	for i, s := range m {
		for _, value := range s {
			key := strings.ToLower(value)
			result[string(key)] += i
		}
	}
	return result
}
