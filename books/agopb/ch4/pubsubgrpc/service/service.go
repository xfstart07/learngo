// Author: xufei
// Date: 2019-08-26

package service

import (
	"context"
	"fmt"
	"learngo/books/agopb/ch1/pubsub"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (s *PubsubService) Publish(ctx context.Context, args *String) (*String, error) {
	fmt.Println("发送信息")
	s.pub.Publish(args.GetValue())
	return &String{}, nil
}

func (s *PubsubService) Subscribe(args *String, stream PubsubService_SubscribeServer) error {
	ch := s.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			fmt.Println("订阅信息", v, key)
			if strings.HasPrefix(key, args.GetValue()) {
				return true
			}
		}
		fmt.Println("无效消息", v)
		return false
	})

	for v := range ch {
		if err := stream.Send(&String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}
