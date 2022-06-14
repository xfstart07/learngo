// Author: Xu Fei
// Date: 2018/8/17
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i:= 0; i < len(os.Args); i++ {
		// 这种字符串拼接的效率比较低
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
