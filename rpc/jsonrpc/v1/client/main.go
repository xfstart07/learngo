// Author: xufei
// Date: 2018/12/17

package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"

	rpcdemo "learngo/rpc/jsonrpc/v1"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)
	defer client.Close() // 关闭连接

	args := rpcdemo.Args{
		A: 10,
		B: 2,
	}
	var result float64
	err = client.Call("DivisionDemo.Div", args, &result)
	if err != nil {
		fmt.Printf("err %v", err)
	} else {
		fmt.Println(result)
	}

	args = rpcdemo.Args{
		A: 20,
		B: 3,
	}
	err = client.Call("DivisionDemo.Div", args, &result)
	if err != nil {
		fmt.Printf("err %v", err)
	} else {
		fmt.Println(result)
	}
}
