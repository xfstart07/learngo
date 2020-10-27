// Author: xufei
// Date: 2019-06-04

package main

import "fmt"

func main() {
	fmt.Println("first.fmt.Println")
	defer fmt.Println("first.defer.fmt.Println")

	fmt.Println("second.fmt.Println")
	defer fmt.Println("second.defer.fmt.Println")
}

//first.fmt.Println
//second.fmt.Println
//second.defer.fmt.Println
//first.defer.fmt.Println

// 可以看出来，defer 的调用顺序是一个栈的形式，先申请的后调用
