// Author: xufei
// Date: 2019-08-23

package service

import (
	"context"
	"fmt"
)

type HelloServiceImpl struct{}

func (s *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	fmt.Println("receive hello...")
	reply := &String{
		Value: "hello:" + args.GetValue(),
	}

	return reply, nil
}
