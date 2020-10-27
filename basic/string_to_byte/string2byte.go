// Author: xufei
// Date: 2019-12-30

package main

import (
	"fmt"
	"unsafe"
)

// 零拷贝将 string 转为 []byte
// 通过将指针地址转为 unsafe.Pointer，在转为另一个类型的指针实现
// 因为 string 和 []byte 的底层数据结构体都是一个整型数组
func main() {
	// []byte to string
	bytes := []byte{104, 101, 108, 108, 111}
	p := unsafe.Pointer(&bytes)
	str := (*string)(p)
	fmt.Println(*str)

	// string to []byte
	pb := unsafe.Pointer(str)
	bytes2 := (*[]byte)(pb)
	fmt.Println(*bytes2)

	fmt.Println(uintptr(0))  // 0
	fmt.Println(^uintptr(0)) // 无穷大的意思
}
