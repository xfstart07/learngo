package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4}
	printValue(arr...)
}

func printValue(arr ...int) {
	fmt.Println(arr)
}
