package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)
var ErrInvalidNumber = errors.New("invalid phone number")

func Number(phoneNumber string) (string, error) {
	re := regexp.MustCompile("[0-9]+")
	found := re.FindAllString(phoneNumber, -1)
	numbers := strings.Join(found, "")
	if len(numbers) == 11 && numbers[0] == '1' {
		numbers = numbers[1:]
	}
	if len(numbers) != 10 || numbers[0] == '0' || numbers[3] == '0' || numbers[0] == '1' || numbers[3] == '1' {
		return "", ErrInvalidNumber
	}
	return numbers, nil
}

func AreaCode(phoneNumber string) (string, error) {
	if number, err := Number(phoneNumber); err == nil {
		return number[:3], nil
	} else {
		return "",  ErrInvalidNumber
	}
}

func Format(phoneNumber string) (string, error) {
	if number, err := Number(phoneNumber); err == nil {
		return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
	} else {
		return "",  ErrInvalidNumber
	}
}
