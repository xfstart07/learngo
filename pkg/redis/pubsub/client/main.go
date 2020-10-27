// Author: xufei
// Date: 2020/5/26

package main

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

// 订阅者
// 只有一个订阅者执行，加入reids 分布式锁
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.4.71:6379",
	})

	log.Println(rdb.Ping().Result())

	pubsub := rdb.Subscribe("task1")

	ch := pubsub.Channel()

	rdb.Del("task_lock")
	for msg := range ch {
		lock, err := rdb.SetNX("task_lock", "lock", 24*time.Hour).Result()
		log.Println(lock, err)
		if lock {
			log.Println(msg.Channel, msg.Payload)
			rdb.Del("task_lock")
		}

	}
}
