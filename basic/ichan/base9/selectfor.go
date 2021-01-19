// Author: xufei
// Date: 2021/1/19

package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)

	syncChan := make(chan struct{}, 1)
	go func() {
	Loop:
		for {
			select {
			case e, ok := <-intChan:
				if !ok {
					fmt.Println("End.")
					break Loop // NOTE: 中断紧贴 Loop 标签之下的 for 语句
				}
				fmt.Printf("Received: %v\n", e)
			}
		}
		syncChan <- struct{}{}
	}()
	<-syncChan
}
