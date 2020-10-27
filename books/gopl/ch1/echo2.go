// Author: Xu Fei
// Date: 2018/8/17
package main

import (
	"os"
	"fmt"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		// 这种字符串拼接的效率比较低
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
