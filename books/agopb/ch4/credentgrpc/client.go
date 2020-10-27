// Author: xufei
// Date: 2019-08-28

package main

import (
	"context"
	"fmt"
	"learngo/books/agopb/ch4/hellogrpc/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	crets, err := credentials.NewClientTLSFromFile("agopb/ch4/credentgrpc/server.crt", "server.grpc.io")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("weixinote.dev:5000", grpc.WithTransportCredentials(crets))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := service.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &service.String{Value: "world"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())

}
