// Author: xufei
// Date: 2019-05-15

package main

import (
	"fmt"
	"learngo/gobasics/funcational/options/app"
)

func main() {
	// 将操作通过函数传入
	c := app.New(app.UserAgent("chrome"))
	fmt.Println(c)
}
