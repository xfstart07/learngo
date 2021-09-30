// Author: xufei
// Date: 2021/1/19

package main

import "fmt"

func main() {
	chanCap := 5
	intChan := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		//NOTE: 多个 case 的选择是随机性的
		select {
		case intChan <- 3:
		case intChan <- 1:
		case intChan <- 2:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%d\n", <-intChan)
	}
}
