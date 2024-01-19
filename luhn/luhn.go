package luhn

import "strconv"

func Valid(id string) bool {
	digits := make([]int, 0, len(id))
	for _, ch := range id {
		if n, ok := strconv.Atoi(string(ch)); ok == nil {
			digits = append(digits, n)
		} else if ch == ' ' {
			continue
		} else {
			return false
		}
	}
	if len(digits) < 2 {
		return false
	}
	for i := len(digits) - 2; i >= 0; i -= 2 {
		digits[i] *= 2
		if digits[i] > 9 {
			digits[i] -= 9
		}
	}
	sum := 0
	for _, n := range digits {
		sum += n
	}
	return sum%10 == 0
}
