// Author: xufei
// Date: 2019-12-24

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "weixinote.dev:6379",
		Password: "",
		DB:       0,
	})

	var lockKey = "counter_lock"
	var counterKey = "counter"

	//	lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockResult, err := resp.Result()
	if err != nil || !lockResult {
		fmt.Println(err, "lock result:", lockResult)
		return
	}

	//	counter++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			fmt.Println(err, "increment counter error!")
		}
	}
	fmt.Println("current counter is", cntValue)

	delResp := client.Del(lockKey)
	unLockResult, err := delResp.Result()
	if err != nil || unLockResult == 0 {
		fmt.Println("unlock failed", err)
	}
	fmt.Println("unlock success!")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}

	wg.Wait()
}
