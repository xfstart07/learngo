// Author: Xu Fei
// Date: 2018/8/23
package main

import (
	"time"
	"fmt"
)

func main() {
	go spinner(100*time.Microsecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}


func fib(x int) int{
	if x < 2 {
		return x
	}

	return fib(x-1)+fib(x-2)
}