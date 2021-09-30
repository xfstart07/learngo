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
	go receive(strChan, syncChan1, syncChan2)
	go send(strChan, syncChan1, syncChan2)
	<-syncChan2
	<-syncChan2
}

func receive(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {
	<-syncChan1
	fmt.Println("Received a sync signal and wait a second...[receiver]")
	time.Sleep(time.Second)
	//for 语句会不断地尝试从通道中接收元素值， 直到该通道关闭。 在通道关闭时， 如果通道中已无元素值， 那么这条 for 语句的执行就会立即结束。
	for elem := range strChan {
		fmt.Println("Received:", elem, "[receiver]")
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
	//close(strChan) // Cannot use 'strChan' (type <-chan string) as type chan<- Type
}

func send(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync signal. [sender]")
		}
	}
	fmt.Println("Wait 2 second... [sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2 <- struct{}{}
}
