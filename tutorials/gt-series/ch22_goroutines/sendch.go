// Author: Xu Fei
// Date: 2018/7/27
package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan<- int)
	go sendData(sendch)
	fmt.Println(<-sendch) // 无效：因为 sendch 是一个只输入的 channel，不支持读取
	//	运行会打印错误 invalid operation: <-sendch (receive from send-only type chan<- int)
}
