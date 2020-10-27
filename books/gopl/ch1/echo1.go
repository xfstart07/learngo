// Author: Xu Fei
// Date: 2018/8/17
package main

import (
	"os"
	"fmt"
)

func main() {
	var s, sep string
	for i:= 1; i < len(os.Args); i++ {
		// 这种字符串拼接的效率比较低
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// 字符串拼接的基准测试在 echo_test 中