// Author: xufei
// Date: 2019-08-22

package main

import (
	"fmt"
	"learngo/books/agopb/ch4/jsonrpc/service"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	// jsonrpc 包实现了 Dial 方法，里面的创建过程同下
	//conn, err := net.Dial("tcp", "weixinote.dev:1234")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	client, err := jsonrpc.Dial("tcp", "weixinote.dev:1234")
	if err != nil {
		log.Fatal(err)
	}

	var reply string
	err = client.Call(service.HelloServiceName+".Hello", "security service", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
