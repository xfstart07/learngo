// Author: xufei
// Date: 2019-10-23

package main

import "fmt"

func main() {
	deferCall()
}

func deferCall() {
	defer func() {
		fmt.Println("打印前")
		/*if err := recover(); err != nil {
		    fmt.Println(err)
		}*/
	}()

	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}
