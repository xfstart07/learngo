// Author: xufei
// Date: 2019-08-22

package main

import (
	"fmt"
	v1 "learngo/books/agopb/ch4/protobuf/v1"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "weixinote.dev:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply v1.String
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err) // gob: type mismatch in decoder: want struct type v1.String; got non-struct
	}

	fmt.Println(reply)
}
