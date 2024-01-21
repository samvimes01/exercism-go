package isbn

import "strconv"

func IsValidISBN(isbn string) bool {
	if len(isbn) < 10 {
		return false
	}
	sum := 0

	if isbn[len(isbn)-1] == 'X' {
		sum += 10
	} else {
		if last, err := strconv.Atoi(string(isbn[len(isbn) - 1])); err != nil {
			return false
		} else {
			sum += last
		}
	}
	pos := 0
	for _, ch := range isbn[:len(isbn)-1] {
		if ch == '-' {
			continue
		}
		if digit, err := strconv.Atoi(string(ch)); err != nil {
			return false
		} else {
			sum += digit * (10 - pos)
		}
		pos++
	}
	return sum%11 == 0
}
