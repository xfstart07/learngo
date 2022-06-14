// Author: Xu Fei
// Date: 2018/8/17
package main

import (
	"os"
	"fmt"
)

func main() {
	for n, arg := range os.Args[1:] {
		fmt.Println(n, arg)
	}
}
