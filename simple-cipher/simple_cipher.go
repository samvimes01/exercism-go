package cipher

import (
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift int
type vigenere string

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance > 25 || distance == 0 || distance < -25 {
		return nil
	}
	return shift(distance)
}

func (c shift) Encode(input string) string {
	builder := strings.Builder{}
	for _, l := range input {
		if unicode.IsLetter(l) {
			ch := unicode.ToLower(l)
			ciph := ((ch + 26 + rune(c) - 97) % 26) + 97
			builder.WriteRune(ciph)
		}
	}
	return builder.String()
}

func (c shift) Decode(input string) string {
	builder := strings.Builder{}
	for _, ch := range input {
		ciph := ((ch + 26 - rune(c) - 97) % 26) + 97
		builder.WriteRune(ciph)
	}
	return builder.String()
}

func NewVigenere(key string) Cipher {
	allA := true
	for _, b := range []byte(key) {
		if b < 'a' || b > 'z' {
			return nil
		}
		allA = allA && b == 'a'
	}

	if allA {
		return nil
	}
	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	return v.translate(input, 1)
}

func (v vigenere) Decode(input string) string {
	return v.translate(input, -1)
}

func (v vigenere) translate(in string, dir int) string {
	var out strings.Builder
	out.Grow(len(in))
	for _, b := range []byte(in) {
		n := int(b | ' ' - 'a') // lowercase letter position 0 - 26
		if n >= 0 && n <= 26 {
			shiftN := int(v[out.Len()%len(v)]-'a')
			ciph := (n + dir*shiftN + 26) %26 + 'a'
			out.WriteByte(byte(ciph))
		}
	}
	return out.String()

}
