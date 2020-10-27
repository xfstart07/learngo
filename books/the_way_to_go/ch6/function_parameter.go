package main

import (
	"fmt"
)

func main() {
	callback(2, Add)
}

func Add(a, b int) int {
	return a + b
}

func callback(y int, f func(int, int) int) {
	fmt.Println(f(y, 2))
}
