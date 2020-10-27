// Author: xufei
// Date: 2019-08-23

package main

import (
	"context"
	"fmt"
	"learngo/books/agopb/ch4/tokengrpc/service"
	"log"

	"google.golang.org/grpc"
)

func main() {
	auth := service.Authentication{
		User:     "gopher",
		Password: "password",
	}

	conn, err := grpc.Dial("weixinote.dev:1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}

	client := service.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &service.String{Value: "world, for auth"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())
}

//hello:world
