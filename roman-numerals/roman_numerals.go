package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("input out of range")
	}
	str := strings.Builder{}

	for input > 0 {
		if input >= 1000 {
			str.WriteString("M")
			input -= 1000
			continue
		}
		if input >= 900 {
			str.WriteString("CM")
			input -= 900
			continue
		}
		if input >= 500 {
			str.WriteString("D")
			input -= 500
			continue
		}
		if input >= 400 {
			str.WriteString("CD")
			input -= 400
			continue
		}
		if input >= 100 {
			str.WriteString("C")
			input -= 100
			continue
		}
		if input >= 90 {
			str.WriteString("XC")
			input -= 90
			continue
		}
		if input >= 50 {
			str.WriteString("L")
			input -= 50
			continue
		}
		if input >= 40 {
			str.WriteString("XL")
			input -= 40
			continue
		}
		if input >= 10 {
			str.WriteString("X")
			input -= 10
			continue
		}
		if input >= 9 {
			str.WriteString("IX")
			input -= 9
			continue
		}
		if input >= 5 {
			str.WriteString("V")
			input -= 5
			continue
		}
		if input >= 4 {
			str.WriteString("IV")
			input -= 4
			continue
		}
		if input >= 1 {
			str.WriteString("I")
			input -= 1
			continue
		}
	}
	return str.String(), nil
}
