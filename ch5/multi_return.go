package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	div, remainder, err := divAndReminder(10, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(div, remainder)
}

// mutliple return values: types are listed in Parana thesis
// you must return all of the values, separated by commas but without paranthesis around them
func divAndReminder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}
