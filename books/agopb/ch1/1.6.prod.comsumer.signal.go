// Author: xufei
// Date: 2019-08-20

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Comsumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

// 生产者消费者模型
func main() {
	ch := make(chan int, 64)

	go Producer(3, ch)
	go Producer(5, ch)

	go Comsumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
