// Author: xufei
// Date: 2019-08-22

package main

import (
	v1 "learngo/books/agopb/ch4/protobuf/v1"
	"log"
	"net"
	"net/rpc"
)

func main() {
	err := rpc.RegisterName("HelloService", new(v1.HelloService))
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
