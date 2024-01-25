package allyourbase

import "errors"

func ConvertToBase(base int, digits []int, outputBase int) (conversion []int, err error) {
	if base < 2 {
		return []int{0}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{0}, errors.New("output base must be >= 2")
	}
	var sum int
	for _, d := range digits {
		if d >= base || d < 0 {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		sum = sum*base + d
	}
	if sum == 0 {
		return []int{0}, nil
	}
	for ; sum > 0; sum /= outputBase {
		conversion = append([]int{sum % outputBase}, conversion...)
	}
	return conversion, nil
}
/*
https://www.tutorialspoint.com/computer_logical_organization/number_system_conversion.htm
Step 1 − Divide the decimal number to be converted by the value of the new base.

Step 2 − Get the remainder from Step 1 as the rightmost digit (least significant digit) of new base number.

Step 3 − Divide the quotient of the previous divide by the new base.

Step 4 − Record the remainder from Step 3 as the next digit (to the left) of the new base number.

Repeat Steps 3 and 4, getting remainders from right to left, until the quotient becomes zero in Step 3.
*/
