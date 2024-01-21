package main

import "fmt"

func main() {
	// C-style
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// just like if statement, there i s no paranthesis around the parts of the for statements
	// there are three parts separated by a columns
	//
	// 1. initialization you set one or more variable before the loop starts
	// 	use := to initialzie, var is not allowed here, you can shadow a variable in the if statements
	// 2. comparision: this must be an expressiona that evaluated to a bool
	// 3. increment, (any assigment is valid here)
	// the condition only for statement
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
	// the infinite for loop
	// for {
	// 	fmt.Println("Hello")
	// }
	// the for-range statement
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	// if you don't need the i, you can use _
	fmt.Println("-------------------")
	fmt.Println("not using the index")
	for _, v := range evenVals {
		fmt.Println(v)
	}

	// go allows you only using the first value and leaving off the second value
	// this could be useful when iterating over maps that are used as a set
	uniqueName := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueName {
		fmt.Println("name: ", k)
	}
	// when iterating over maps, you can not count on orders
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 10; i++ {
		fmt.Println("New loop:")
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	// the order of keys and values might varies
	fmt.Println("-------------------")
	fmt.Println("Iterating over strings")
	samples := []string{"hello", "apple_خ!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}
	// The for range value is a copy
}
