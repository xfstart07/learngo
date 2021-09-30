// Author: xufei
// Date: 2019-08-15

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	iptr := &i

	fmt.Println("iptr", iptr)

	var fptr *float64 = (*float64)(unsafe.Pointer(iptr))

	*fptr = *fptr * 3

	fmt.Println("fptr", fptr)
	fmt.Println("iptr", iptr)
}
