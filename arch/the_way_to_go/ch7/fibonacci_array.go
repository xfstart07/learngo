package main

import (
	"fmt"
)

func main() {
	var fn [51]uint64

	fn[0] = 0
	fn[1] = 1
	for i := 2; i <= 50; i++ {
		fn[i] = fn[i-2] + fn[i-1]
	}

	for i := range fn {
		fmt.Println(fn[i])
	}

}
