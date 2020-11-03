// Author: xufei
// Date: 2020/11/3

package zset

import (
	"fmt"
	"testing"

	"learngo/pkg/goredis/client"

	"github.com/go-redis/redis"
)

func TestZSet(t *testing.T) {
	cli := client.Init()
	fmt.Println("有序集合")

	fmt.Println("添加集合")
	cnt := cli.ZAdd("pageRank", redis.Z{Score: 10, Member: "google.com"}, redis.Z{Score: 9, Member: "baidu.com"}, redis.Z{Score: 7, Member: "bing.com"}).Val()
	fmt.Println("结果: ", cnt)

	fmt.Println("查询元素值")
	val := cli.ZScore("pageRank", "google.com").Val()
	fmt.Println("结果: ", val)

	fmt.Println("增加权重")
	val = cli.ZIncr("pageRank", redis.Z{Score: 1, Member: "bing.com"}).Val()
	fmt.Println("结果: ", val)

	fmt.Println("统计个数")
	cnt = cli.ZCount("pageRank", "9", "10").Val()
	fmt.Println("结果：", cnt)
}
