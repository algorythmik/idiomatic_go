package main

import "fmt"

func main() {
	twoFactor := makeMulti(2)
	threeFactor := makeMulti(3)
	for i := 0; i < 5; i++ {
		fmt.Println(i, twoFactor(i), threeFactor(i))
	}

}

func makeMulti(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
