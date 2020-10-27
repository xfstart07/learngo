// Author: xufei
// Date: 2019-08-22

package service

import (
	"fmt"
)

// 服务方法接口
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

// 服务具体类型
type HelloService struct{}

func (s *HelloService) Hello(req string, reply *string) error {
	fmt.Println("hello request...")
	*reply = "hello: " + req
	return nil
}
