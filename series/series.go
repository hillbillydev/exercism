package series

func All(n int, s string) []string {
	return chunkString(n, s)
}

func UnsafeFirst(n int, s string) string {
	return chunkString(n, s)[0]
}

func First(n int, s string) (first string, ok bool) {
	c := chunkString(n, s)
	if len(c) < 1 {
		return "", false
	}
	return c[0], true
}

func chunkString(n int, s string) []string {
	var result []string

	for i := range s {
		if len(s) < i+n {
			break
		}
		result = append(result, s[i:i+n])
	}

	return result
}
