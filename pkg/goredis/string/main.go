// Author: xufei
// Date: 2020/11/3

package main

import (
	"fmt"
	"time"

	"learngo/pkg/goredis/client"
)

// 数据类型：字符串
func main() {
	cli := client.Init()

	fmt.Println("存储命令: set")
	if err := cli.Set("redis", "redis.com", 10*time.Second).Err(); err != nil {
		panic(err)
	}

	fmt.Println("获取命令：get")
	val := cli.Get("redis").Val()
	fmt.Println("获取的值", val)

	fmt.Println("存储锁命令：setnx")
	nx := cli.SetNX("redis", "redis.org", 10*time.Second).Val()
	fmt.Println("存储结果: ", nx)

	nx = cli.SetNX("job", "golang", 10*time.Second).Val()
	fmt.Println("存储结果: ", nx)

	fmt.Println("批量存储命令: mset")
	set := cli.MSet("date", "2020.11.3", "time", "11:25", "weater", "tuesday").Err()
	fmt.Println("批量存储结果: ", set)

	vals := cli.MGet("date", "time").Val()
	fmt.Println("批量获取结果", vals)
}

// 还有处理 Integer 类型的 Incr，Decr
