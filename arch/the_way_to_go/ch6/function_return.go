package main

import (
	"fmt"
)

func main() {
	a := add()
	fmt.Println("a的值", a(2))

	a2 := adder(2)
	fmt.Println("a2的值", a2(3))
}

func add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func adder(a int) func(b int) int {
	return func(b int) int {
		return b + a
	}
}
