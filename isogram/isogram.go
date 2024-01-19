package isogram

import "strings"

func IsIsogram(word string) bool {
	chrs := make(map[string]struct{})
	item := struct{}{}
	for _, r := range word {
		ch := strings.ToLower(string(r))
		if _, ok := chrs[ch]; ok {
			return false
		}
		if ch == "-" || ch == " " {
			continue
		}
		chrs[ch] = item
	}
	return true
}
