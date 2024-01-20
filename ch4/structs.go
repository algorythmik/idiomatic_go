package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}
	var fred person
	fmt.Println("fred", fred)
	fmt.Printf("name: %s, age: %d, pet: %s\n", fred.name, fred.age, fred.pet)
	// struct literal
	bob := person{}
	fmt.Printf("bob: %s\n", bob)
	julia := person{
		"Julia",
		23,
		"cat"}
	fmt.Println(julia)
	beth := person{
		age:  39,
		name: "Beth"}
	fmt.Println("Beth", beth)
	fmt.Println(beth.name, beth.age, beth.pet)
}
