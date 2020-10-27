// Author: xufei
// Date: 2019-08-28

package main

import (
	"fmt"
	"learngo/books/agopb/ch4/hellogrpc/service"
	"log"
	"net/http"
	"strings"

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

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "hello")
	})

	// grpc 需要建立在 HTTP2 上，否则无法使用
	log.Fatal(http.ListenAndServeTLS(":5000", "agopb/ch4/credentgrpc/server.crt", "agopb/ch4/credentgrpc/server.key", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor != 2 {
			mux.ServeHTTP(w, r)
			return
		}

		if strings.Contains(
			r.Header.Get("Context-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
			return
		}

		mux.ServeHTTP(w, r)
		return
	})))
}
