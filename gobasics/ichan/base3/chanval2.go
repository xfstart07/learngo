// Author: xufei
// Date: 2021/1/18

package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

func (c *Counter) String() string {
	return fmt.Sprintf("{count: %d}", c.count)
}

var mapChan = make(chan map[string]*Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() {
		countMap := map[string]*Counter{
			"count": &Counter{},
		}
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
// var mapChan = make(chan map[string]Counter, 1)
//The count map: map[count:{0}]. [sender]
//The count map: map[count:{0}]. [sender]
//The count map: map[count:{0}]. [sender]
//The count map: map[count:{0}]. [sender]
//The count map: map[count:{0}]. [sender]
// 因为 go 传递是值拷贝，所有接收者的操作对发送者不会有影响

// var mapChan = make(chan map[string]*Counter, 1)
//The count map: map[count:{count: 1}]. [sender]
//The count map: map[count:{count: 2}]. [sender]
//The count map: map[count:{count: 3}]. [sender]
//The count map: map[count:{count: 4}]. [sender]
//The count map: map[count:{count: 5}]. [sender]
// 这里 mapChan 是引用类型，所有接收者的操作对发送者也会有影响
