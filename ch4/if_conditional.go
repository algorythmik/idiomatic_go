package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That is too low")
	} else if n <= 5 {
		fmt.Println("That is a good number")
	} else {
		fmt.Println("That is too high")
	}
	fmt.Println("n: ", n)
	// you can declare variables that are scoped to the condition
	if m := rand.Intn(10); m == 0 {
		fmt.Println("That is too low")
	} else {
		fmt.Println("That is too high or reasonable")
	}
	fmt.Println("m: ", m)

}
