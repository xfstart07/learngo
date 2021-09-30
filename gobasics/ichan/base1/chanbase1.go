// Author: xufei
// Date: 2021/1/18

package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	// 演示接收操作
	go func() {
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second... [receiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Printf("Received: %v, [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	// 演示发送操作
	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent: ", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync signal. [sender]")
			}
		}
		fmt.Println("Wait 2 seconds...[sender]")
		time.Sleep(2 * time.Second)
		close(strChan)
		syncChan2 <- struct{}{}
	}()

	<-syncChan2
	<-syncChan2
}
