package lsproduct

import (
	"errors"
	"math"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var max int64 = math.MinInt64
	if span > len(digits) {
		return 0, errors.New("span cannot be greater than digits")
	}
	if span < 0 {
		return 0, errors.New("span cannot be negative")
	}
	for i := 0; i < len(digits); i++ {
		if i + span > len(digits) {
			break
		}
		spanned := digits[i : i+span]
		product := int64(1)
		for _, d := range spanned {
			d, err := strconv.Atoi(string(d))
			if err != nil {
				return 0, err
			}
			product *= int64(d)
		}
		if product > max {
			max = product
		}
	}
	return max, nil
}
