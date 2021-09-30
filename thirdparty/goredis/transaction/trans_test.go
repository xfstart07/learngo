// Author: xufei
// Date: 2020/11/3

package transaction

import (
	"fmt"
	"testing"
	"time"

	"learngo/thirdparty/goredis/client"

	"github.com/go-redis/redis"
)

func TestTrans(t *testing.T) {
	cli := client.Init()
	fmt.Println("事务")

	pipe := cli.TxPipeline()
	pipe.ZAdd("pageRank", redis.Z{Score: 1, Member: "douban.com"}).Val()
	pipe.Expire("pageRank", 10*time.Second)
	exec, err := pipe.Exec()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("结果", exec)
}
