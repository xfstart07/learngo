package main

import (
	"fmt"
)

func main() {
	x := min(1, 3, 2, 4)
	fmt.Println("最小的数：", x)

	a := []int{1, 2, 34, 3, 5, 6}
	fmt.Println(min(a...))
}

func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if min > v {
			min = v
		}
	}
	return min
}
