package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter2 int64
	wg2      sync.WaitGroup
)

func main() {
	// 在一个处理器下运行
	runtime.GOMAXPROCS(1)

	wg2.Add(2)

	go incCounter2(1)
	go incCounter2(2)

	wg2.Wait()
	fmt.Println("Final Counter:", counter2) // counter = 2
}

func incCounter2(id int) {
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter2, 1)
		// 当前 goroutine 从线程退出，并放回队列
		runtime.Gosched()
	}
}
