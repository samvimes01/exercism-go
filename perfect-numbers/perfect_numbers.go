package perfect

import "errors"

// Define the Classification type here.
type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("only positive")

// Define the functions here.
func SumFactors(n int64) int64 {
	var sum int64
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	sum := SumFactors(n)
	var class Classification
	switch {
	case n == sum:
		return ClassificationPerfect, nil
	case n < sum:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
	return class, nil
}
