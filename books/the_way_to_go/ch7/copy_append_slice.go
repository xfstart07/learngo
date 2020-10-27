package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3}
	sl2 := make([]int, 10)
	s3 := copy(sl2, sl) // return len
	fmt.Println(s3)
	fmt.Println(sl2)

	s4 := append(sl2, 1, 2)
	s4 = append(sl2, s3...)
	fmt.Println(s4) // [1 2 3 0 0 0 0 0 0 0 1 2]
}
