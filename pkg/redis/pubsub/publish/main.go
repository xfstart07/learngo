// Author: xufei
// Date: 2020/5/26

package main

import (
	"log"

	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.4.71:6379",
	})

	log.Println(rdb.Ping().Result())

	rdb.Publish("task1", "hello")
}
