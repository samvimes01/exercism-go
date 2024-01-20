package pangram

import (
	"strings"
	"unicode"
)

func IsPangram1(input string) bool {
	countMap := make(map[rune]int, 26)
	for _, char := range "abcdefghijklmnopqrstuvwxyz" {
		countMap[char] = 0
	}
	for _, char := range input {
		char = unicode.ToLower(char)
		if char < 'a' || char > 'z' {
			continue
		}
		countMap[char]++
	}
	for _, count := range countMap {
		if count == 0 {
			return false
		}
	}
	return true
}

func IsPangram2(s string) bool {
	lookup := strings.ToLower(s)
	for chr := 'a'; chr <= 'z'; chr++ {
		if !strings.ContainsRune(lookup, chr) {
			return false
		}
	}
	return true
}

const alphamask int = 0b11111111111111111111111111

func IsPangram(phrase string) bool {
	phrasemask := 0
	length := len(phrase)
	for i := 0; i < length; i++ {
		letter := phrase[i]
		if letter > 96 && letter < 123 {
			phrasemask |= 1 << (letter - 97)
		} else if letter > 64 && letter < 91 {
			phrasemask |= 1 << (letter - 65)
		}
	}
	return phrasemask == alphamask
}
