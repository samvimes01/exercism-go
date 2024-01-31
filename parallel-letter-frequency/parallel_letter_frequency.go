package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(strings []string) FreqMap {
	m := FreqMap{}
	results := make(chan FreqMap)
	for _, s := range strings {
		go func(s string) {
			results <- Frequency(s)
		}(s)
	}
	for range strings {
		for k, v := range <-results {
			m[k] += v
		}
	}
	return m
}
