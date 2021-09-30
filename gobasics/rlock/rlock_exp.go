// Author: xufei
// Date: 2019-10-16 16:25

package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

var mlock sync.RWMutex
var wg sync.WaitGroup

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go gets()
	}

	wg.Wait()
}

func gets() {
	for i := 0; i < 100000; i++ {
		get(i)
	}
	wg.Done()
}

func get(i int) {
	beginTime := time.Now()
	mlock.RLock()
	tmp1 := time.Since(beginTime).Nanoseconds() / 1000000
	if tmp1 > 100 { // 超过100ms就打印出来
		fmt.Println("fuck here")
	}
	mlock.RUnlock()
}

// 运行：
// go run back2.go 2> trace

// 查看 trace 信息
// go tool trace ./trace
// 会打开一个浏览器网页，View trace 能看到 goroutine 的运行情况，
// Goroutine analysis 内的 main.gets 能看到所有 goroutine 详细的运行时间。

// https://xargin.com/a-rlock-story/
// 这篇文章的例子。学习了 trace 的使用方法。还有锁的使用场景：在锁的范围内使用 syscall 是耗时挺多的。
