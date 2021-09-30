// Author: xufei
// Date: 2020/3/24

package main

import (
	"encoding/base64"
	"fmt"
)

func Append(sli, data []byte) []byte {
	l := len(sli)
	if l+len(data) > cap(sli) {
		// 扩容
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, sli)
		sli = newSlice
	}

	for i, v := range data {
		sli[l+i] = v
	}

	return sli
}

func main() {
	sli := []byte("1fdsfasdf")
	data := []byte("dddddd")

	sli = Append(sli, data)
	fmt.Println(sli, string(sli))

	s, _ := base64.StdEncoding.DecodeString("eGlhbzE3eXk=")
	fmt.Println(string(s))
}
