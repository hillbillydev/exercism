package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
// FreqMap is a map that holds the rune as the key and number of occurences as the value.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
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

	result := FreqMap{}
	temp := []FreqMap{}
	mtx := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, v := range s {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			m := Frequency(text)
			mtx.Lock()
			defer mtx.Unlock()
			temp = append(temp, m)
		}(v)
	}

	wg.Wait()

	for _, t := range temp {
		for k, v := range t {
			result[k] += v
		}
	}

	return result
}
