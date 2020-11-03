// Author: xufei
// Date: 2020/11/3

package hash

import (
	"fmt"
	"testing"

	"learngo/pkg/goredis/client"
)

func TestHASH(t *testing.T) {
	cli := client.Init()
	fmt.Println("哈希表")

	fmt.Println("存储命令: hset")
	ok := cli.HSet("songs", "001", "后来").Val()
	fmt.Println("存储结果", ok)
	cli.HSet("songs", "002", "十年")

	fmt.Println("存储锁命令 hsetnx")
	ok = cli.HSetNX("songs", "003", "成都").Val()
	fmt.Println("存储结果", ok)

	val := cli.HGet("songs", "001").Val()
	fmt.Println("获取内容", val)

	fmt.Println("删除 hdel")
	dv := cli.HDel("songs", "003").Val()
	fmt.Println("结果", dv)
}

// 批量存储，获取
// HMSet, HMGet

// 获取所有的域，值
// HKeys，HVals

func TestHashGet(t *testing.T) {
	cli := client.Init()

	fmt.Println("获取所有信息")
	m := cli.HGetAll("songs").Val()
	fmt.Println("结果: ", m)
}
