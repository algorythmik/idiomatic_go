package main

import (
	"fmt"
	"sort"
)

func main() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"Mojtab", "Tabaei", 42},
		{"Ali", "Reza", 12},
	}

	fmt.Println("Before sort: ", people)
	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("After sort: ", people)

}
