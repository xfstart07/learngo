// Author: xufei
// Date: 2020/11/4

package distributelock

import (
	"sync"
	"testing"
	"time"

	"learngo/thirdparty/goredis/client"
)

func TestLock(t *testing.T) {
	cli := client.Init()
	cli.Del("task1")

	sy := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		sy.Add(1)
		go func(i int) {
			lock := NewLock(cli, "task1", 1*time.Minute)
			ok, err := lock.Lock()
			if err != nil {
				t.Logf("错误 %v", err)
			}
			if ok {
				t.Logf("成功获取到锁: %v", i)
			} else {
				t.Logf("获取失败: %v", i)
			}
			sy.Done()
		}(i)
	}
	// 只会有一个获得锁

	sy.Wait()
}
