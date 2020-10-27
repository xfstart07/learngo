// Author: Xu Fei
// Date: 2018/8/20
package main

import (
	"fmt"
	"bytes"
)

func comma(s string) string{
	buf := bytes.Buffer{}
	if len(s) <= 3 {
		buf.WriteString(s)
	} else {
		for i := range s {
			if (i+1)%3 == 0 {
				buf.WriteByte(',')
			}
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}

// 3.10
func main() {
	s := "1234543423234324"
	s2 := comma(s)
	fmt.Println(s2)
}
