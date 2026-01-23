package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	numberOfCows  int
	customMessage string
}

// TODO: define the 'DivideFood' function
func DivideFood(fodder FodderCalculator, numberOfCows int) (float64, error) {
	totalFodder, err := fodder.FodderAmount(numberOfCows)
	if err != nil {
		return 0, err
	}
	fattener, err := fodder.FatteningFactor()
	if err != nil {
		return 0, err
	}
	foodPerCow := (totalFodder * fattener) / float64(numberOfCows)

	return foodPerCow, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodder FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows > 0 {
		return DivideFood(fodder, numberOfCows)
	}

	return 0, errors.New("invalid number of cows")
}

// TODO: define the 'ValidateNumberOfCows' function
func (er *InvalidCowsError) Error() string {
	if er.numberOfCows < 0 {
		return fmt.Sprintf("%d cows are invalid: there are no negative cows", er.numberOfCows)
	}
	if er.numberOfCows == 0 {
		return fmt.Sprintf("%d cows are invalid: no cows don't need food", er.numberOfCows)
	}
	return ""
}

func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows > 0 {
		return  nil
	}
	return &InvalidCowsError{
		numberOfCows: numberOfCows,
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
