// Author: xufei
// Date: 2020/3/5

package main

import (
	"context"
	"learngo/gomicro/second-part/prime-srv/handler"
	LogProto "learngo/gomicro/second-part/proto/log"
	"learngo/gomicro/second-part/proto/prime"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.learning.srv.prime"),
		micro.Broker(rabbitmq.NewBroker()),
	)

	srv.Init(
		micro.WrapHandler(reLogger(srv.Client())),
	)

	_ = prime.RegisterPrimeHandler(srv.Server(), handler.NewHandler())

	if err := srv.Run(); err != nil {
		panic(err)
	}
}

func reLogger(cli client.Client) server.HandlerWrapper {
	pub := micro.NewEvent("go.micro.learning.topic.log", cli)
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			log.Log("准备发送日志")
			evt := LogProto.LogEvt{Msg: "Call Prime-Srv"}

			err := pub.Publish(context.TODO(), &evt)
			if err != nil {
				log.Errorf("log send failed: %v", err)
			}

			return handlerFunc(ctx, req, rsp)
		}
	}
}
