// Author: xufei
// Date: 2020/11/3

package client

import "github.com/go-redis/redis"

func Init() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "weixinote.dev:6379",
		Password: "",
	})

	return client
}
