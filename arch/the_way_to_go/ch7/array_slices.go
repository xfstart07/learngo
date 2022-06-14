package main

import (
	"fmt"
)

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Println("切片内容", slice1[i])
	}

	fmt.Printf("数组长度%d\n", len(arr1))
	fmt.Printf("切片长度 %d\n", len(slice1))
	fmt.Printf("切片容量 %d\n", cap(slice1))

}
