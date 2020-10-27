// Author: xufei
// Date: 2019-08-22

package service

import (
	"fmt"
	"log"
	"net"
)

// 服务方法接口
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
	Login(req string, reply *string) error
}

// 服务具体类型
type HelloService struct {
	conn    net.Conn
	isLogin bool
}

func NewHelloService(conn net.Conn) *HelloService {
	return &HelloService{
		conn: conn,
	}
}

func (s *HelloService) Login(req string, reply *string) error {
	fmt.Println("login...")

	if req != "user:password" {
		return fmt.Errorf("auth failed")
	}

	log.Println("login ok")
	s.isLogin = true
	*reply = "login ok!"
	return nil
}

func (s *HelloService) Hello(req string, reply *string) error {
	if !s.isLogin {
		return fmt.Errorf("please login")
	}

	fmt.Println("service handle request...")
	*reply = "hello: " + req + ", form: " + s.conn.RemoteAddr().String()
	return nil
}
