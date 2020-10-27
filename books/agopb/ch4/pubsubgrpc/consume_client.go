// Author: xufei
// Date: 2019-08-27

package main

import (
	"context"
	"fmt"
	"io"
	"learngo/books/agopb/ch4/pubsubgrpc/service"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("weixinote.dev:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := service.NewPubsubServiceClient(conn)

	stream, err := client.Subscribe(context.Background(), &service.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				continue
			}
			log.Fatal(err)
		}

		fmt.Println("接收结果", recv.GetValue())
	}

}
