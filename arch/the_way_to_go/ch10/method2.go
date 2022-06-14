package main

import (
	"fmt"
)

type IntVector []int

func main() {
	fmt.Println("和：", IntVector{1, 2, 3}.Sum()) // 6
}

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}
