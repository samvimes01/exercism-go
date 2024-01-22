package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(map[string]int)
	word := make([]rune, 0)
	count := func() {
		if len(word) > 0 && word[0] == '\'' {
			word = word[1:]
		}
		if len(word) > 0 && word[len(word)-1] == '\'' {
			word = word[:len(word)-1]
		}
		if len(word) > 0 {
			k := strings.ToLower(string(word))
			freq[k]++
			word = make([]rune, 0)
		}
	}
	for _, c := range phrase {
		if unicode.IsLetter(c) || unicode.IsDigit(c) || c == '\'' {
			word = append(word, c)
		} else {
			count()
		}
	}
	count()
	return freq
}

/*
// regexp

func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`\w+('\w+)?`)
	count := make(Frequency)
	for _, w := range re.FindAllString(strings.ToLower(phrase), -1) {
		count[w]++
	}
	return count 
}
*/