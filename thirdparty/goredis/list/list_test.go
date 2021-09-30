// Author: xufei
// Date: 2020/11/3

package list

import (
	"fmt"
	"testing"

	"learngo/thirdparty/goredis/client"
)

func TestList(t *testing.T) {
	cli := client.Init()
	fmt.Println("队列命令")

	fmt.Println("插入头部（左）")
	ret := cli.LPush("lang", "python", "go").Val()
	fmt.Println("结果: ", ret)

	fmt.Println("插入尾部（右）")
	ret = cli.RPush("lang", "ruby").Val()
	fmt.Println("结果：", ret)

	fmt.Println("获取队列")
	val := cli.LRange("lang", 0, -1).Val()
	fmt.Println("结果: ", val)

	fmt.Println("移除一个")
	s := cli.LPop("lang").Val()
	fmt.Println("结果:", s)

	fmt.Println("移除某个相同值")
	cnt := cli.LRem("lang", 0, "ruby").Val()
	fmt.Println("移除个数: ", cnt)
}
