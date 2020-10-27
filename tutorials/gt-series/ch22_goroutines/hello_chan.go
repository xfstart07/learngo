// Author: Xu Fei
// Date: 2018/7/27
package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("hello world goroutine")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")
}
