package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pumps(ch2)
	go sucks(ch1, ch2)

	time.Sleep(1e9)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func pumps(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func sucks(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Receive 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received 2: %d\n", v)
		}
	}
}
