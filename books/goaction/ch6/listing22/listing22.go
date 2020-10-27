package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 无缓冲的通道，做同步操作
func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("Runner %d Running with Baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Println("Runner Finished, Race Over. ", runner)
		wg.Done()

		return
	}

	fmt.Printf("Runner %d Exchange with Runner %d.\n", runner, newRunner)

	baton <- newRunner
}
