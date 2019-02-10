// Package twofer makes it easy to share items with your mates.
package twofer

import "fmt"

// ShareWith return a sentence depending of the input.
// If the input name is not specified it uses "you" as a default.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
