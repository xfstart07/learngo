// Author: xufei
// Date: 2019-08-23

package main

import (
	"context"
	"fmt"
	"learngo/books/agopb/ch4/filtergrpc/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("filter:", info)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}

func main() {
	// 构造一个 gRPC 服务对象
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(filter))
	// 注册 gRPC 服务
	service.RegisterHelloServiceServer(grpcServer, service.NewHelloServiceImpl("gopher", "password"))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// 通过 gRPC serve 监听端口连接
	log.Fatal(grpcServer.Serve(listener))
}
