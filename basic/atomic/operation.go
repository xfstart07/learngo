// Author: xufei
// Date: 2019/3/10

package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var add int32
	atomic.AddInt32(&add, 1)
	fmt.Println(add)
}
