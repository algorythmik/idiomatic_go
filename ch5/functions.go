package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(div(10, 2))
	// variadic input parameters and slices
	fmt.Println(addTo(1, 2, 3, 4, 5))

	// multiple return parameters
	fmt.Println(divAndReminder(10, 0))
}

// when you have multiple input parameters of the same type:
func div(numerator, denumerator int) int {
	if denumerator == 0 {
		return 0
	}
	return numerator / denumerator
}

// variadic input funcitons
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// mutliple return values: types are listed in Parana thesis
// you must return all of the values, separated by commas but without paranthesis around them
func divAndReminder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}
