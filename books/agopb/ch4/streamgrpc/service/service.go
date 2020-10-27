// Author: xufei
// Date: 2019-08-23

package service

import (
	"context"
	"fmt"
	"io"
)

type HelloServiceImpl struct{}

func (s *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}
		fmt.Println("args", args.GetValue())

		reply := &String{Value: "hello" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (s *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	fmt.Println("receive hello...")
	reply := &String{
		Value: "hello:" + args.GetValue(),
	}

	return reply, nil
}
