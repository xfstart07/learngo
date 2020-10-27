// Author: Xu Fei
// Date: 2018/8/20
package main

import "fmt"

func main() {
	a := []string{"one", "two", "two","tree"}

	// 不配置 len 的长度，但配置 cap 的容量
	b := make([]string, 0, len(a))
	b = append(b, a[0])
	for i := 1; i< len(a); i++ {
		if a[i-1] == a[i] {
			continue
		} else {
			b = append(b, a[i])
		}
	}
	fmt.Println(b)
}

