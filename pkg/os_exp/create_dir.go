package main

import (
	"os"
	"fmt"
	"syscall"
)

func main() {
	filePath := "public/kk"

	_, err := os.Stat(filePath)

	b := err == nil || os.IsExist(err)
	fmt.Println("是否存在", b)

	if !b {
		mask := syscall.Umask(0)
		defer syscall.Umask(mask)
		err := os.MkdirAll(filePath, 0666)
		if err != nil {
			fmt.Println("创建文件夹失败", err)
		}
	}

}
