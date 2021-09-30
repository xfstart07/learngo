// Author: xufei
// Date: 2019-08-15

package main

import (
	"fmt"
	"unsafe"
)

type user1 struct {
	b byte
	i int32
	j int64
}

func main() {
	var u1 user1

	fmt.Println(unsafe.Offsetof(u1.b))
	fmt.Println(unsafe.Offsetof(u1.i))
	fmt.Println(unsafe.Offsetof(u1.j))
}

//0
//4
//8
