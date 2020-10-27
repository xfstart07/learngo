// Author: xufei
// Date: 2019-08-22

package v1

import "fmt"

type HelloService struct{}

func (s *HelloService) Hello(req *String, reply *String) error {
	fmt.Println("hello request...")
	reply.Value = "hello: " + req.Value
	return nil
}
