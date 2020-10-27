// Author: xufei
// Date: 2018/12/25

package main

import "fmt"

type Slice []int

// 因为是值引用，所有调用 Append 不会改变 mSlice
// https://www.flysnow.org/2018/12/21/golang-sliceheader.html
//
func (A Slice) Append(value int) {
	A1 := append(A, value)
	fmt.Printf("%p\n%p\n", A, A1)
	fmt.Printf("len = %d, %d\n", len(A), len(A1))
	fmt.Printf("cap = %d %d\n", cap(A), cap(A1))
	A = A1
}

func main() {
	//mSlice := make(Slice, 10, 20)
	mSlice := make(Slice, 10, 10)
	fmt.Printf("mSlice = %p\n", mSlice)
	mSlice.Append(5)
	fmt.Println(mSlice) // [0 0 0 0 0 0 0 0 0 0]

	a := make([]int, 10, 10)
	b := append(a, 5)
	fmt.Printf("%p\n%p\n", a, b)
}
