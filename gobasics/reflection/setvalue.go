// Author: xufei
// Date: 2019-08-19

package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(3)
	fmt.Println(x)
}
