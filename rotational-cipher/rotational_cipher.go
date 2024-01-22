package rotationalcipher

import "strings"

func RotationalCipher(plain string, shiftKey int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	rotated := alphabet[shiftKey:] + alphabet[:shiftKey]
	mapped := make(map[string]string, 52)
	for i, letter := range alphabet {
		mapped[string(letter)] = string(rotated[i])
		mapped[strings.ToUpper(string(letter))] = string(strings.ToUpper(string(rotated[i])))
	}
	cipher := strings.Builder{}
	for _, letter := range plain {
		if ch, ok := mapped[string(letter)]; ok {
			cipher.WriteString(ch)
		} else {
			cipher.WriteRune(letter)
		}
	}
	return cipher.String()
}
