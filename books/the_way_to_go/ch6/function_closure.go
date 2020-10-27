package main

import (
	"fmt"
)

func main() {
	var f = adderc()
	fmt.Println("1 返回:", f(1))
	fmt.Println("21 返回:", f(21))
	fmt.Println("212 返回:", f(212))
}

// 当 adderc 调用后，会返回一个闭包，这时每次调用 x 的值都会增加
func adderc() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
