// Author: Xu Fei
// Date: 2018/8/20
package main

import (
	"fmt"
	"bytes"
)

func intsToString(values []int) string{
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {

	fmt.Println(intsToString([]int{1,2,3,4,5,6}))
}
