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
func ConcurrentFrequency(texts []string) FreqMap {
	frequencies := FreqMap{}
	done := make(chan struct{}, len(texts))
	for _, text := range texts {
		ch := parseRunes(text)
		for r := range ch {
			frequencies[r]++
		}
		done <- struct{}{}
	}
	for i:=0; i<len(texts); i++ {
		<-done
	}
	return frequencies
}

func parseRunes(text string) <-chan rune {
	ch := make(chan rune)
	go func() {
		for _, r := range text {
			ch <- r
		}
		close(ch)
	}()
	return ch
}
