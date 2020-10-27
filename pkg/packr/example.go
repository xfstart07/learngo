package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/packr"
)

func main() {
	// templates 是相对路径，例子是在同一个目录下
	box := packr.NewBox("./templates")

	// 以字符串的形式获取静态文件
	html := box.String("index.html")
	fmt.Println("String获取:", html)

	html, err := box.MustString("index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MustString获取", html)

	// 以字节数组的形式获取静态文件
	htmlByte := box.Bytes("index.html")
	fmt.Println("Bytes: ", htmlByte)
	// 对应的还有 MustBytes 方法
}
