// Author: xufei
// Date: 2019-08-12

package rwmutex

import (
	"fmt"
	"sync"
	"time"
)

type hash struct {
	m  map[string]string
	mu sync.RWMutex
}

func UseRead(h *hash) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("获取读锁...")
			h.mu.RLock()
			time.Sleep(10 * time.Millisecond)
			fmt.Println(h.m["name"]) // leon
			h.mu.RUnlock()
		}()
	}
}

func UseWrite(h *hash, value string) {
	fmt.Println("获取写锁...")
	h.mu.Lock()
	h.m["name"] = value
	h.mu.Unlock()
}

func Use() {
	h := &hash{
		m: make(map[string]string),
	}

	UseWrite(h, "leon")
	UseRead(h)
	time.Sleep(10 * time.Millisecond)
	UseWrite(h, "kevin") // 这次获取写锁时会阻塞，需要等所有读锁释放

	time.Sleep(1 * time.Second)
}
