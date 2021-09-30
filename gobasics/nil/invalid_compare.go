// Author: xufei
// Date: 2019-10-24

package main

import (
	"fmt"
)

func main() {
	var v1 []int = nil
	var v2 []int = nil
	fmt.Println(v1 == v2) // invalid operation: v1 == v2 (operator ==  not defined on []int)
	fmt.Println((map[string]int)(nil) == (map[string]int)(nil))
	fmt.Println((func())(nil) == (func())(nil))
}
