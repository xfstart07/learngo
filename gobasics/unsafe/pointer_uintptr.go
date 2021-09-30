// Author: xufei
// Date: 2019-08-15

package main

import (
	"fmt"
	"unsafe"
)

type user struct {
	name string
	age  int
}

func main() {
	u := new(user)
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	*pName = "张三"

	pAge := (*int)(unsafe.Pointer(&u.age))
	//pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 20

	fmt.Println(*u)
}

//{ 0}
//{张三 20}
