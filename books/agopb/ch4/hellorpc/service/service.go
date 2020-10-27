// Author: xufei
// Date: 2019-08-22

package service

import "fmt"

type HelloService struct{}

func (s *HelloService) Hello(req string, reply *string) error {
	fmt.Println("hello request...")
	*reply = "hello: " + req
	return nil
}
