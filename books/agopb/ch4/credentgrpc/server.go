// Author: xufei
// Date: 2019-08-28

package main

import (
	"learngo/books/agopb/ch4/hellogrpc/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cret, err := credentials.NewServerTLSFromFile("agopb/ch4/credentgrpc/server.crt", "agopb/ch4/credentgrpc//server.key")
	if err != nil {
		log.Fatal(err)
	}

	// 构造一个 gRPC 服务对象
	grpcServer := grpc.NewServer(grpc.Creds(cret))

	// 注册 gRPC 服务
	service.RegisterHelloServiceServer(grpcServer, new(service.HelloServiceImpl))

	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	// 通过 gRPC serve 监听端口连接
	log.Println("server run...")
	log.Fatal(grpcServer.Serve(listener))
}
