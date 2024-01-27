package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	div, remainder, err := divAndReminder(10, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(div, remainder)
	// you can ignore a value explicitly
	div, _, err = divAndReminder(10, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(div)
	// or you can impliciltly ignore ALL of the values
	divAndReminder(10, 2)

}

// mutliple return values: types are listed in Parana thesis
// you must return all of the values, separated by commas but without paranthesis around them
func divAndReminder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	// unlike python, in go these are actually multiple values and not a tuple
	return numerator / denominator, numerator % denominator, nil
}

// you can have named return values
func nameDivAndReminder(numerator, denominator int) (result int, remainder int, err error) {

	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	// unlike python, in go these are actually multiple values and not a tuple
	return numerator / denominator, numerator % denominator, nil
}
