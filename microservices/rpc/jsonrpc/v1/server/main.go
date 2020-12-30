// Author: xufei
// Date: 2018/12/17

package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	rpcdemo "learngo/microservices/rpc/jsonrpc/v1"
)

func main() {
	err := rpc.Register(rpcdemo.DivisionDemo{})
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
