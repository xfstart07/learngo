// Author: Xu Fei
// Date: 2018/8/8
package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "weixinote.dev:6379",
		Password: "",
		DB:       0,
	})

	client.Set("hello", "world", -1)
	fmt.Println(client.Get("hello").String())

	fmt.Println(client.Exists("hello").Err())
	fmt.Println(client.Exists("hello").Result())

	if cmd := client.Del("hello"); cmd.Err() != nil {
		fmt.Println(cmd.Err())
	}

	fmt.Println(client.Exists("hello").Err())
	fmt.Println(client.Exists("hello").Result())
}
