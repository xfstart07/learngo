// Author: Xu Fei
// Date: 2018/3/27
package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5}

	pos := 2
	fmt.Println(a[pos:])
	fmt.Println(a[pos+1:])
	fmt.Println(copy(a[pos:], a[pos+1:]))
	fmt.Println(a)
	//a = a[:pos+copy(a[pos:], a[pos+1:])]

	fmt.Println(a)

}