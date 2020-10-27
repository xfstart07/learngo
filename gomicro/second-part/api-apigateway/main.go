// Author: xufei
// Date: 2020/3/6

package main

import (
	"context"
	"encoding/json"
	"learngo/gomicro/second-part/proto/prime"
	"learngo/gomicro/second-part/proto/sum"
	"log"
	"strconv"

	api "github.com/micro/go-micro/v2/api/proto"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"

	proto "learngo/gomicro/second-part/proto/api"
)

var (
	sumClient   sum.SumService
	primeClient prime.PrimeService
)

type Open struct {
}

func (o *Open) Fetch(ctx context.Context, req *api.Request, resp *api.Response) error {
	// 打印请求头
	for k, v := range req.Header {
		log.Println("请求头信息，", k, " : ", v)
	}
	log.Println("----------------------------")

	sumInputStr := req.Get["sum"].Values[0]
	primeInputStr := req.Get["prime"].Values[0]

	sumInput, _ := strconv.ParseInt(sumInputStr, 10, 64)
	sumReq := &sum.SumRequest{
		Input: sumInput,
	}
	primeInput, _ := strconv.ParseInt(primeInputStr, 10, 64)
	primeReq := &prime.PrimeRequest{
		Input: primeInput,
	}

	sumResp, _ := sumClient.GetSum(ctx, sumReq)
	primeResp, _ := primeClient.GetPrime(ctx, primeReq)

	ret, _ := json.Marshal(map[string]interface{}{
		"sum":   sumResp.Output,
		"prime": primeResp.Output,
	})

	resp.Body = string(ret)

	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.learning.api.open"),
	)

	srv.Init(micro.Action(func(c *cli.Context) error {
		sumClient = sum.NewSumService("go.micro.learning.srv.sum", srv.Client())
		primeClient = prime.NewPrimeService("go.micro.learning.srv.prime", srv.Client())
		return nil
	}))

	// 暴露接口
	_ = proto.RegisterOpenHandler(srv.Server(), new(Open))

	if err := srv.Run(); err != nil {
		log.Fatal("api gateway running fail", err)
	}
}
