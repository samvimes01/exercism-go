package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	cnt := 0
	if n <= 0 {
		return 0, errors.New("invalid input, should be bigger 0")
	}
	for n > 1 {
		cnt++
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return cnt, nil
}
