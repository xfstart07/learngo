// Author: xufei
// Date: 2019-08-23

package main

import (
	"learngo/books/agopb/ch4/pubsubgrpc/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 构造一个 gRPC 服务对象
	grpcServer := grpc.NewServer()
	// 注册 gRPC 服务
	service.RegisterPubsubServiceServer(grpcServer, service.NewPubsubService())

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// 通过 gRPC serve 监听端口连接
	log.Fatal(grpcServer.Serve(listener))
}
