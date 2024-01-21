package romannumerals

import (
	"errors"
	"strings"
)

var decRomMap = []struct {
	decimal int
	roman   string
}{
	{1, "I"}, {4, "IV"}, {5, "V"}, {9, "IX"},
	{10, "X"}, {40, "XL"}, {50, "L"}, {90, "XC"},
	{100, "C"}, {400, "CD"}, {500, "D"}, {900, "CM"},
	{1000, "M"},
}

func ToRomanNumeral(dec int) (string, error) {
	if dec < 1 || dec > 3999 {
		return "", errors.New("number must be in range 1-3999")
	}
	roman := ""
	repeatRoman := 0
	for i := len(decRomMap) - 1; i >= 0; i-- {
		decimal := decRomMap[i].decimal
		repeatRoman, dec = dec/decimal, dec%decimal
		roman += strings.Repeat(decRomMap[i].roman, repeatRoman)
	}
	return roman, nil
}

func ToRomanNumeral2(input int) (string, error) {
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
