// Author: xufei
// Date: 2020/11/3

package set

import (
	"fmt"
	"testing"

	"learngo/thirdparty/goredis/client"
)

func TestSet(t *testing.T) {
	cli := client.Init()
	fmt.Println("集合")

	fmt.Println("存储集合")
	val := cli.SAdd("website", "baidu.com", "github.com", "discuz.net").Val()
	fmt.Println("结果: ", val)
	cli.SAdd("bbs", "discuz.net")

	fmt.Println("判断是否在集合中")
	ok := cli.SIsMember("website", "github.com").Val()
	fmt.Println("结果: ", ok)

	fmt.Println("移除元素")
	val = cli.SRem("website", "baidu.com").Val()
	fmt.Println("结果:", val)

	fmt.Println("交集")
	strs := cli.SInter("website", "bbs").Val()
	fmt.Println("结果: ", strs)

	fmt.Println("并集")
	strs = cli.SUnion("website", "bbs").Val()
	fmt.Println("结果: ", strs)

	fmt.Println("差集")
	strs = cli.SDiff("website", "bbs").Val()
	fmt.Println("结果: ", strs)
}
