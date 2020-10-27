// Author: xufei
// Date: 2019-12-24

package main

import (
	"log"
	"sync"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	log.SetFlags(log.Lshortfile)

	conn, _, err := zk.Connect([]string{"weixinote.dev"}, time.Second)
	if err != nil {
		panic(err)
	}
	log.Println("connect state:", conn.State().String())

	l := zk.NewLock(conn, "/lock", zk.WorldACL(zk.PermAll))

	var counter int
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = l.Lock() // 会阻塞
			if err != nil {
				log.Println("get lock failed", err)
				return
			}
			log.Println("get lock success!")

			counter++
			log.Println("current counter:", counter)

			err := l.Unlock()
			log.Println("unlock!", err)
		}()
	}

	wg.Wait()
}
