package armstrong

import "math"

func IsNumber(n int) bool {
	digits := getDigits(n)
	pow := len(digits)
	sum := 0
	for _, digit := range digits {
		sum += int(math.Pow(float64(digit), float64(pow)))
	}
	return sum == n
}

func getDigits(n int) []int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return digits
}