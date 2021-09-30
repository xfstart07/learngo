// Author: xufei
// Date: 2019-10-15

package main

import (
	"log"
	"sync"
	"sync/atomic"

	"learngo/thirdparty/mq/nsq/publish"
)

func main() {
	var songID int32 = 0
	var errCount int32 = 0

	wg := sync.WaitGroup{}

	for i := 0; i < 30; i++ {
		go func() {
			wg.Add(1)

			if err := publish.NewProduct(&songID); err != nil {
				atomic.AddInt32(&errCount, 1)

				log.Println(err)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	log.Println("发送成功！", "发送错误: ", errCount, songID)
}
