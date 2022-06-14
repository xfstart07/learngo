// 10.6.3

package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func main() {
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)\n", lst, lst.Len()) // [1] (len: 1)

	ptrList := new(List)
	ptrList.Append(2)
	fmt.Printf("%v (len %d)\n", ptrList, ptrList.Len()) // &[2] (len 1)
}
