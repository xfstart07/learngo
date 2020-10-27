// 这个示例程序展示如何在程序里造成竞争状态

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// 共享资源
	counter int
	wg      sync.WaitGroup
)

func main() {
	// 在一个处理器下运行
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter) // counter = 2
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		// 当前 goroutine 从线程退出，并放回队列
		runtime.Gosched()
		value++
		counter = value
	}
}
