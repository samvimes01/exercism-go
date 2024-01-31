package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	plain := "abcdefghijklmnopqrstuvwxyz"
	cipher := "zyxwvutsrqponmlkjihgfedcba"
	dict := make(map[rune]rune)
	for i, v := range plain {
		dict[v] = rune(cipher[i])
	}
	ciphered := []rune{}
	for _, v := range s {
		if l, ok := dict[unicode.ToLower(v)]; ok {
			ciphered = append(ciphered, l)
		} else if unicode.IsLetter(v) || unicode.IsNumber(v) {
			ciphered = append(ciphered, v)
		}
	}
	builder := strings.Builder{}
	for i, v := range ciphered {
		builder.WriteRune(v)
		if (i+1)%5 == 0 {
			builder.WriteRune(' ')
		}
	}
	return strings.TrimSpace(builder.String())
}
