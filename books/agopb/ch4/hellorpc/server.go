// Author: xufei
// Date: 2019-08-22

package main

import (
	"learngo/books/agopb/ch4/hellorpc/service"
	"log"
	"net"
	"net/rpc"
)

func main() {
	err := rpc.RegisterName("HelloService", new(service.HelloService))
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error ", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error", err)
	}

	rpc.ServeConn(conn)
}
