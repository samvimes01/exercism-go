package reverse

import "strings"

func Reverse(input string) string {
	str := strings.Builder{}
	asRune := []rune(input)
	for i:= len(asRune)-1; i>=0; i-- {
		str.WriteRune(asRune[i])
	}
	return str.String()
}
