// Author: xufei
// Date: 2021/1/12

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGQUIT)
	fmt.Printf("接收的信号: %s\n", <-sig)
}
