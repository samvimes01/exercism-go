package acronym

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Abbreviate(s string) string {
	re := regexp.MustCompile(`[a-zA-Z']+`) // a-z cos \w == a-z_
	words := re.FindAllString(s, -1)
	abbr := make([]rune, 0)
	for _, word := range words {
		ch, size := utf8.DecodeRuneInString(word)
		if size > 0 && unicode.IsLetter(ch) {
			abbr = append(abbr, ch)
		}
	}

	return strings.ToUpper(string(abbr))

}
