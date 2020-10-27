// Author: Xu Fei
// Date: 2018/8/20
package main

import "fmt"

func basename(s string) string {
	// 去掉尾部的 / 以上的部分
	for i := len(s) -1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s)-1; i>=0; i-- {
		if s[i]== '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println("abc")
}