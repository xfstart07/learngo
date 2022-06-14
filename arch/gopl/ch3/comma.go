// Author: Xu Fei
// Date: 2018/8/20
package main

import "fmt"

func comma(s string) string{
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	s := "12345"
	s2 := comma(s)
	fmt.Println(s2)
}
