package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		// shadowing: x out of the current block is shadhowed
		// and x in the  next line in the current block is the shadowing variable
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

// if you just remove : from the line 11, x is not shadowed but is reassigned a value

// func main() {
// 	x := 10
// 	if x > 5 {
// 		fmt.Println(x)
// 		// shadowing: x out of the current block is shadhowed
// 		// and x in the  next line in the current block is the shadowing variable
// 		x = 5
// 		fmt.Println(x)
// 	}
//  fmt.Println(x)
//}
// with := it is easy to accidentally shadow a variable.
