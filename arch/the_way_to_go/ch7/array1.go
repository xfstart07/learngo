package main

import (
	"fmt"
)

func main() {
	a := [...]string{"a", "b", "c"}
	for i := range a {
		fmt.Println(a[i])
	}
	// abc

	a1 := a // 值拷贝
	fmt.Printf("a 的内存 %p\n", &a)
	fmt.Printf("a1 的内存 %p\n", &a1)

	var arrLazy = [...]int{5, 6, 7, 8, 22} // 这是一个切片
	fmt.Println("cap", cap(arrLazy))

	var arr2 [5]int
	fmt.Println("cap", cap(arr2))

}
