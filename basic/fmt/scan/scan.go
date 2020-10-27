// Author: Leon Xu
// Date: 9 Sep 2019

package main

import (
	"fmt"
	"log"
)

func main() {
	for {
		var text string
		log.Print("请输入: ")
		_, err := fmt.Scanf("%s\n", &text)
		if err != nil {
			log.Fatal("读取错误", err)
		}

		log.Println("用户输入信息: ", text)
	}
}
