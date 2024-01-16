package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalc FodderCalculator, numCows int) (float64, error) {
  amnt, err := fodderCalc.FodderAmount(numCows)
  if err != nil {
    return 0.0, err
  }
  fc, err := fodderCalc.FatteningFactor()
  if err != nil {
    return 0.0, err
  }
  return amnt * fc / float64(numCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalc FodderCalculator, numCows int) (float64, error) {
  if numCows > 0 {
    return DivideFood(fodderCalc, numCows)
  } else {
    return 0.0, errors.New("invalid number of cows")
  }
}

type InvalidCowsError struct {
  cows int
  msg string
}
func (e *InvalidCowsError) Error() string {
  return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.msg)
}
// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error {
  if numCows < 0 {
    return &InvalidCowsError{numCows, "there are no negative cows"}
  } else if numCows == 0 {
    return &InvalidCowsError{numCows, "no cows don't need food"}
  } else {
    return nil
  }
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
