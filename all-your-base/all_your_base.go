package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	if len(inputDigits) == 0 || len(inputDigits) == 1 && inputDigits[0] == 0 {
		return []int{0}, nil
	}
	// convert input base to base 10
	digits10 := 0
	for i, v := range inputDigits {
		if v < 0 || v >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		pos := len(inputDigits) - i - 1
		digits10 += v * int(math.Pow(float64(inputBase), float64(pos)))
	}
	if digits10 == 0 {
		return []int{0}, nil
	}
	// convert to output base
	outputDigits := make([]int, 0)
	for digits10 > 0 {
		digitBase := digits10 % outputBase
		digits10 /= outputBase
		outputDigits = append(outputDigits, digitBase)
	}
	// reverse
	for i, j := 0, len(outputDigits)-1; i < j; i, j = i+1, j-1 {
		outputDigits[i], outputDigits[j] = outputDigits[j], outputDigits[i] 
	}
	return outputDigits, nil
}
/*
https://www.tutorialspoint.com/computer_logical_organization/number_system_conversion.htm
Step 1 − Divide the decimal number to be converted by the value of the new base.

Step 2 − Get the remainder from Step 1 as the rightmost digit (least significant digit) of new base number.

Step 3 − Divide the quotient of the previous divide by the new base.

Step 4 − Record the remainder from Step 3 as the next digit (to the left) of the new base number.

Repeat Steps 3 and 4, getting remainders from right to left, until the quotient becomes zero in Step 3.
*/
