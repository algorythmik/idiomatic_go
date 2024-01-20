package main

import "fmt"

func main() {
	// unlike arrays in slices length is not part of the type of the slice
	// slices are somewhat similar to arrays but there are some subtle differences.
	// we don't specify the size of a slice when we declare the slice
	// declaration using a slice literal
	x := []int{10, 20, 30}
	fmt.Println(x)
	// this creates a slice of 3 ints using slice literal
	// if you have sparse slice you can use the followings
	y := []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("y:", y)
	// we can create multidimensional slices:
	var z [][]int
	fmt.Println("z", z)
	// zero value for a slice:
	var zero_val_slice []int
	fmt.Println(zero_val_slice == nil)
	fmt.Println("zero_val_slice:", zero_val_slice)
	fmt.Println("len built in function")
	fmt.Println("you can use len built-in functiont o get the length of slice")
	fmt.Println("length of zero_val_slice:", len(x))

	// make
	// you can use make when you want to create a slice with capacity specified.
	make_slice := make([]int, 5)
	fmt.Println("make_slice: ", make_slice)
	fmt.Println("Empty literal slice declaration")
	emty_literal_slice := []int{}
	fmt.Println("empty_literal_slice=", emty_literal_slice)
	// this creates a slice that  is not nil
	fmt.Println(emty_literal_slice == nil)
	// but with length of zero
	fmt.Println(len(emty_literal_slice))
	// declaring slice with default value
	data := []int{2, 4, 6, 8}
	fmt.Println("data ", data)
	// slicing slices
	m := []int{1, 2, 3, 4, 5}
	n := m[:2]
	fmt.Println("m, n", m, n)
	fmt.Printf("m %s, n: %s", m, n)
}
