package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	fmt.Println(add1(a, b))
	fmt.Println(add2(a, b))
}

func add1(a, b int) int {
	return a + b
}

func add2(a, b int) (sum int) {
	sum = a + b
	return
}
