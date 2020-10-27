// Author: Xu Fei
// Date: 2018/8/24
package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	//time.Sleep(2*time.Second) // 无法打印
	ch <- 1
	close(ch)
}

func read(ch chan int) {
	fmt.Println(<-ch) // 多次读，在读取了写入的值之后，之前读取到的是零值
}

func main() {
	ch := make(chan int)
	go write(ch)
	go read(ch) //1
	go read(ch) //0
	time.Sleep(1*time.Second)
}