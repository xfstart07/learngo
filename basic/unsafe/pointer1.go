// Author: xufei
// Date: 2019-08-15

package main

import "fmt"

func main() {
	i := 10
	iptr := &i

	fmt.Println(iptr)

	//var fptr *float64 = (*float64)(iptr)
	//cannot convert expression of type *int to type *float64

	//fmt.Println(fptr)
}
