package main

import "fmt"

func main() {
	a := []float32{1.1, 2.2, 3.3}
	fmt.Println(sumDouble(a))

	// 问题 7.7

	fmt.Println(a[1:1]) // 空数组
	fmt.Println(a[0:1])
}

func sumDouble(a []float32) float32 {
	var sum float32
	for ix := range a {
		sum += a[ix]
	}
	return sum
}
