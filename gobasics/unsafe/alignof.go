// Author: xufei
// Date: 2019-08-15

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool
	var i8 int8
	var i16 int16
	var i64 int64

	var f32 float32

	var s string

	var m map[string]string

	var p *int32

	var ch chan int

	fmt.Println(unsafe.Alignof(b))   // 1个字节
	fmt.Println(unsafe.Alignof(i8))  //  1个字节
	fmt.Println(unsafe.Alignof(i16)) // 2 个字节
	fmt.Println(unsafe.Alignof(i64)) // 8 个字节
	fmt.Println(unsafe.Alignof(f32)) // 4 个字节
	fmt.Println(unsafe.Alignof(s))   // 8个字节
	fmt.Println(unsafe.Alignof(m))   // 8 个字节
	fmt.Println(unsafe.Alignof(p))   // 指针(unsafe.Pointer) 8 个字节
	fmt.Println(unsafe.Alignof(ch))  // 8个字节
}

//1
//1
//2
//8
//4
//8
//8
//8
//8
