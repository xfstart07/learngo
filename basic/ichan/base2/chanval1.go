// Author: xufei
// Date: 2021/1/18

package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	// 接收
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stoped. [receiver]")
		syncChan <- struct{}{}
	}()
	// 发送
	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Microsecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

// Output
//The count map: map[count:1]. [sender]
//The count map: map[count:2]. [sender]
//The count map: map[count:3]. [sender]
//The count map: map[count:3]. [sender]
//The count map: map[count:5]. [sender]
//Stoped. [receiver]
