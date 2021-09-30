// Author: xufei
// Date: 2021/1/19

package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
	select {
	case e := <-intChan:
		fmt.Printf("Received: %v\n", e)
	case <-time.After(time.Millisecond * 500):
		fmt.Println("timeout!")
	}
}
