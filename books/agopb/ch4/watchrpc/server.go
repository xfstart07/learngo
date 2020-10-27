// Author: xufei
// Date: 2019-08-22

package main

import (
	"learngo/books/agopb/ch4/watchrpc/service"
	"log"
	"net"
	"net/rpc"
)

func main() {
	err := service.RegisterService(service.NewKVStoreService())
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal()
		}

		rpc.ServeConn(conn)
	}
}
