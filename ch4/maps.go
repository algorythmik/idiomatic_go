package main

import (
	"fmt"
)

func main() {
	// var declaration
	var nilMap map[string]int
	fmt.Println("nilMap", nilMap)
	// using an empty map literal
	totalWins := map[string]int{}
	fmt.Println(totalWins)
	teams := map[string][]string{
		"Orcans":  []string{"fred", "ralph", "Bijou"},
		"Lions":   []string{"sarah", "peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	for k, v := range teams {
		fmt.Printf("members of team %s are %s\n", k, v)
	}
	// reading and writing a map
	fmt.Printf("value for team is %s.\n", teams["Orcans"])
	// the comma ok idiom
	m := map[string]int{
		"hello": 5,
		"wolrd": 0,
		"!":     0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["goodbye"]
	fmt.Println(v, ok)
	// deleting from maps
	delete(m, "!")
	fmt.Println(m)
	inSet := map[int]bool{}
	vals := []int{5, 10, 2, 5}
	for _, v := range vals {
		inSet[v] = true
	}
	fmt.Println("len vals:", len(vals))
	if inSet[10] {
		fmt.Println("10 is in the set")
	}
	fmt.Println(inSet[100])
}
