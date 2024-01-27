package main

import (
	"fmt"
	"strconv"
)

func main() {
	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
	expressions := [][]string{
		{"1", "+", "2"},
		{"1", "-", "2"},
		{"1", "*", "2"},
		{"1", "/", "2"},
		{"1", "^", "2"},
		{"1", "^"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression")
			continue
		}
		op, ok := opMap[expression[1]]
		if !ok {
			fmt.Println("operator not supperted!")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := op(p1, p2)
		fmt.Println(result)
	}
}
func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }
