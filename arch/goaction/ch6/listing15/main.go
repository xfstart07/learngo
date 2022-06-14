package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

// 使用原子函数进行同步访问资源操作，避免资源竞争
func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	go dowork("A")
	go dowork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Shutdown now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

func dowork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s working\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
