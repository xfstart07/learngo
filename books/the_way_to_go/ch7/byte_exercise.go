package main

import (
	"fmt"
)

func main() {
	s1 := []byte("12311ff")
	data := []byte("dde")

	fmt.Println(appendByte(s1, data))
}

func appendByte(slice, data []byte) []byte {
	sliceLen := len(slice)
	if sliceLen+len(data) > cap(slice) {
		newSlice := make([]byte, (sliceLen+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : sliceLen+len(data)]
	for i, c := range data {
		slice[sliceLen+i] = c
	}
	return slice
}
