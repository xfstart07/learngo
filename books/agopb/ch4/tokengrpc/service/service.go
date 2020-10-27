// Author: xufei
// Date: 2019-08-23

package service

import (
	"context"
	"fmt"
)

type HelloServiceImpl struct {
	auth *Authentication
}

func NewHelloServiceImpl(user, password string) *HelloServiceImpl {
	return &HelloServiceImpl{
		auth: &Authentication{
			User:     user,
			Password: password,
		},
	}
}

func (s *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	if err := s.auth.Auth(ctx); err != nil {
		return nil, err
	}

	fmt.Println("receive hello...")
	reply := &String{
		Value: "hello:" + args.GetValue(),
	}

	return reply, nil
}
