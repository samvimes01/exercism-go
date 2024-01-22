package rotationalcipher

import "strings"

func RotationalCipher(plain string, shiftKey int) string {
	cipher := strings.Builder{}
	for _, letter := range plain {
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
			shift := letter + rune(shiftKey)
			if shift > 'z' && letter <= 'z' || shift > 'Z' && letter <= 'Z' {
				shift -= 26
			}
			cipher.WriteRune(shift)
		} else {
			cipher.WriteRune(letter)
		}
	}
	return cipher.String()
}
