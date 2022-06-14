package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump2(ch1) // pump hangs
	go suck2(ch1)
	time.Sleep(1e9)
	// fmt.Println(<-ch1) // prints only 0
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck2(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
