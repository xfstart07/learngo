// Author: xufei
// Date: 2019-08-22

package service

import (
	"fmt"
	"net/rpc"
)

// 服务名称
const HelloServiceName = "jsonrpc/HelloService"

// 服务方法接口
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

// 服务注册
func RegisterHelloService(srv HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, srv)
}

// 服务具体类型
type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	fmt.Println("hello request...")
	*reply = "hello: " + request
	return nil
}
