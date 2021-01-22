// Author: xufei
// Date: 2020/3/5

package main

import (
	"context"
	LogProto "learngo/gomicro/second-part/proto/log"
	"learngo/gomicro/second-part/proto/sum"
	"learngo/gomicro/second-part/sum-srv/handler"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

//go run main.go --broker=rabbitmq
func main() {
	// 创建服务
	service := micro.NewService(
		micro.Name("go.micro.learning.srv.sum"),
		micro.Broker(rabbitmq.NewBroker()),
	)

	// 结合Wrapper与Broker
	// 初始化
	service.Init(
		micro.WrapHandler(rateLimit(10)),
		micro.WrapHandler(
			reqLogger(service.Client()),
		),
		micro.BeforeStart(func() error {
			log.Log("[sum] 启动前...")
			return nil
		}),
		micro.AfterStart(func() error {
			log.Log("[sum] 结束前...")
			return nil
		}),
	)

	// 挂载接口
	_ = sum.RegisterSumHandler(service.Server(), handler.NewHandler())

	// 启动
	if err := service.Run(); err != nil {
		panic(err)
	}
}

// 限流
// 限制并发访问，当并发访问到一定数量，limit 就会阻塞，等待之前的请求处理完才能处理新的请求
func rateLimit(rate int) server.HandlerWrapper {
	limit := make(chan bool, rate)
	for i := 0; i < rate; i++ {
		limit <- true
	}

	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			k := <-limit
			defer func() {
				limit <- k
			}()

			return handlerFunc(ctx, req, rsp)
		}
	}
}

func reqLogger(cli client.Client) server.HandlerWrapper {
	// 创建 publisher
	pub := micro.NewEvent("go.micro.learning.topic.log", cli)

	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			log.Info("准备发送日志")
			evt := LogProto.LogEvt{Msg: "Call Sum-Srv"}

			err := pub.Publish(context.TODO(), &evt)
			if err != nil {
				log.Errorf("log send failed: %v", err)
			}

			return handlerFunc(ctx, req, rsp)
		}
	}
}
