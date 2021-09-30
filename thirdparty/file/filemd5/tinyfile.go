// Author: Xu Fei
// Date: 2018/9/13
package main

import (
	"fmt"
	"crypto/md5"
	"io"
	"os"
)

func main() {
	filePath := "/Users/x/golang/src/example/Readme.md"

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	md5h := md5.New()
	io.Copy(md5h, file)
	fmt.Printf("%x", md5h.Sum([]byte(""))) //md5
}
