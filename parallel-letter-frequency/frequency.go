package letter

// FreqMap records the frequency of each rune in a given text.
// FreqMap is a map that holds the rune as the key and number of occurences as the value.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this.
// data as a FreqMap.
// Frequency takes a string and figures out the frequency of those words.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency takes a slice of strings
// and concurrently finds the frequency of those words.
func ConcurrentFrequency(s []string) FreqMap {
	var (
		result  = FreqMap{}
		channel = make(chan FreqMap, 5)
	)
	for _, v := range s {
		go func(text string) {
			channel <- Frequency(text)
		}(v)
	}

	for range s {
		for k, v := range <-channel {
			result[k] += v
		}
	}

	return result
}
