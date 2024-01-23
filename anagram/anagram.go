package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	match :=  make([]string, 0)
	subjectLetters := make(map[rune]int)
	target := strings.ToLower(subject)
	for _, ch := range target {
		subjectLetters[ch]++
	}
	for _, candidate := range candidates {
		hasAllLetters := true
		word := strings.ToLower(candidate)
		
		if word == target {
			continue
		}

		for letter, cnt := range subjectLetters {
			if !strings.ContainsRune(word, letter) || strings.Count(word, string(letter)) < cnt {
				hasAllLetters = false
				break
			}
		}
		if hasAllLetters && len(candidate) == len(subject) {
			match = append(match, candidate)
		}
	}
	return match
}
