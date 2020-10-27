package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go recv(ch1)
	fmt.Println("等10秒")
	time.Sleep(5 * time.Second)
	go send(ch1)

	time.Sleep(30 * time.Second)
}

func recv(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func send(ch chan int) {
	fmt.Println("send")
	ch <- 1
}
