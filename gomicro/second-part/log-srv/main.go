// Author: xufei
// Date: 2020/3/9

package main

import (
	"context"
	proto "learngo/gomicro/second-part/proto/log"

	"github.com/micro/go-plugins/broker/rabbitmq/v2"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

type Sub struct {
}

func (s *Sub) Process(ctx context.Context, evt *proto.LogEvt) error {
	log.Log("[sub] 收到日志: ", evt.Msg)
	return nil
}

//go run main.go --broker=rabbitmq
func main() {
	service := micro.NewService(
		micro.Name("go.micro.learning.srv.log"),
		micro.Broker(rabbitmq.NewBroker()),
	)

	service.Init()

	err := micro.RegisterSubscriber("go.micro.learning.topic.log", service.Server(), &Sub{})
	if err != nil {
		panic(err)
	}

	if err := service.Run(); err != nil {
		panic(err)
	}
}
