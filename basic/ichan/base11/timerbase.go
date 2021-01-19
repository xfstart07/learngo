// Author: xufei
// Date: 2021/1/19

package main

import (
	"fmt"
	"time"
)

// 演示定时器 Timer 和 channel 的使用
func main() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("Present time: %v\n", time.Now())
	expirationTime := <-timer.C
	fmt.Printf("Expiration time: %v\n", expirationTime)
	fmt.Printf("Stop timer: %v", timer.Stop())
}
