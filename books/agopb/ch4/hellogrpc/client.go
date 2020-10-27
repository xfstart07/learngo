// Author: xufei
// Date: 2019-08-23

package main

import (
	"context"
	"fmt"
	"learngo/books/agopb/ch4/hellogrpc/service"
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

	reply, err := client.Hello(context.Background(), &service.String{Value: "world"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())

	// 模拟异步调用
	// 多个Goroutine之间可以安全地共享gRPC底层的HTTP/2链接
	for i := 0; i < 3; i++ {
		go func(idx int) {
			reply, err := client.Hello(context.Background(), &service.String{Value: fmt.Sprintf("world %d", idx)})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(reply.GetValue())
		}(i)
	}

	time.Sleep(1 * time.Second)
}

//hello:world
//hello:world 1
//hello:world 0
//hello:world 2
