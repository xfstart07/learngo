// Author: xufei
// Date: 2021/1/8

package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(10 * time.Second)
	<-timer.C
	fmt.Println("计时器")
}
