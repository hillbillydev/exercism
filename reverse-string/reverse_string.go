package reverse

// String takes a string and reverses it.
func String(s string) string {
	var result string
	for _, rune := range s {
		result = string(rune) + result
	}
	return result
}
