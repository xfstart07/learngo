// Author: xufei
// Date: 2019-08-22

package main

import (
	"learngo/books/agopb/ch4/security_rpc/service"
	"log"
	"net"
	"net/rpc"
)

func main() {
	err := service.RegisterHelloService(&service.HelloService{})
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal()
	}

	rpc.ServeConn(conn)
}
