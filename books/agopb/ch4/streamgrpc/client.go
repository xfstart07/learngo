// Author: xufei
// Date: 2019-08-26

package main

import (
	"context"
	"learngo/books/agopb/ch4/streamgrpc/service"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("weixinote.dev:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := service.NewHelloServiceClient(conn)

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := stream.Send(&service.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}

			time.Sleep(time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}
