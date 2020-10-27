// Author: xufei
// Date: 2019-08-23

package main

import (
	"fmt"
	"learngo/books/agopb/ch4/asynrpc/service"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "weixinote.dev:1234")
	if err != nil {
		log.Fatal(err)
	}

	doClientWork(client)
}

func doClientWork(client *rpc.Client) {
	call := client.Go(service.HelloServiceName+".Hello", "Leon", new(string), nil)

	// do some thing
	fmt.Println("do some things...")

	helloCall := <-call.Done
	if err := helloCall.Error; err != nil {
		log.Fatal(err)
	}

	fmt.Printf("信息 %v \n", helloCall)
	args := helloCall.Args.(string)
	reply := helloCall.Reply.(*string)

	fmt.Println(args, *reply)
}
