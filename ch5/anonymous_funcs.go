package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Printf("print %d from within the function\n", i)
		}(i)
	}
}
