// Author: Xu Fei
// Date: 2018/8/17
package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
