// Author: xufei
// Date: 2019-08-15

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(true))              // bool 占一个字节
	fmt.Println(unsafe.Sizeof(int8(0)))           // int8 类型占一个字节
	fmt.Println(unsafe.Sizeof(int16(10)))         // int16 类型占 2 个字节
	fmt.Println(unsafe.Sizeof(int32(1000000)))    // int32 类型占 4 个字节
	fmt.Println(unsafe.Sizeof(int64(1000000000))) // int64 类型占 8 个字节
	fmt.Println(unsafe.Sizeof(int(100000000000))) // int 类型占 8 个字节，int 是和运行平台有关，在 64位平台是 8个字节，32 位平台是 4 个字节
}

//1
//1
//2
//4
//8
//8
