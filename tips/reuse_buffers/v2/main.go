// Author: xufei
// Date: 2020/1/6

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	pool := make([][]byte, 20)
	buffer := make(chan []byte, 5)

	var m runtime.MemStats
	makes := 0

	for {
		var b []byte

		select {
		// 从缓存中取一个数组
		case b = <-buffer:
		default:
			// 缓存为空，重新生成一个数组
			makes += 1
			b = makeBuffer()
		}

		// 往缓存中放入一个数组，以供下次使用
		i := rand.Intn(len(pool))
		if pool[i] != nil {
			select {
			case buffer <- pool[i]:
				pool[i] = nil
			default:
				//	当缓存已满则跳过，不会阻塞
			}
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
