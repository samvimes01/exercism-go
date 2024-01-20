package pangram

import "unicode"

func IsPangram(input string) bool {
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
