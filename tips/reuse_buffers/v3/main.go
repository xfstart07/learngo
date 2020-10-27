// Author: xufei
// Date: 2020/1/6

package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type queue struct {
	when  time.Time
	slice []byte
}

var (
	makes int
)

func makeBuffer() []byte {
	makes += 1
	return make([]byte, rand.Intn(5000000)+5000000)
}

func makeRecycler() (get, give chan []byte) {
	get = make(chan []byte)
	give = make(chan []byte)

	go func() {
		q := new(list.List)

		for {
			if q.Len() == 0 {
				q.PushFront(queue{when: time.Now(), slice: makeBuffer()})
			}

			item := q.Front()
			timeout := time.NewTimer(time.Minute)
			select {
			case b := <-give:
				timeout.Stop()
				q.PushFront(queue{when: time.Now(), slice: b})
			case get <- item.Value.(queue).slice:
				timeout.Stop()
				q.Remove(item)
			case <-timeout.C:
				e := q.Front()
				for e != nil {
					n := e.Next()
					if time.Since(e.Value.(queue).when) > time.Minute {
						q.Remove(e)
						e.Value = nil
					}

					e = n
				}
			}
		}

	}()

	return
}

func main() {
	pool := make([][]byte, 20)

	get, give := makeRecycler()

	var m runtime.MemStats
	for {
		b := <-get

		// 往缓存中放入一个数组，以供下次使用
		i := rand.Intn(len(pool))
		if pool[i] != nil {
			give <- pool[i]
		}

		pool[i] = b

		time.Sleep(time.Second)

		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("%d,%d,%d,%d,%d,%d\n", makes, m.HeapSys, bytes, m.HeapAlloc, m.HeapIdle, m.HeapReleased)
	}

}
