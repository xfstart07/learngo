// Author: xufei
// Date: 2019-08-27

package main

import (
	"context"
	"learngo/books/agopb/ch4/pubsubgrpc/service"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("weixinote.dev:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := service.NewPubsubServiceClient(conn)
	_, err = client.Publish(context.Background(), &service.String{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(context.Background(), &service.String{Value: "docker: hello Container"})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
}
