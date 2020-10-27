// Author: xufei
// Date: 2019-08-30

package main

import (
	"bufio"
	"fmt"
	"log"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Println(len(p))
	return len(p), nil
}

func main() {
	fmt.Println("无缓冲IO")
	w := new(Writer)
	w.Write([]byte{'a'})
	w.Write([]byte{'b'})
	w.Write([]byte{'c'})
	w.Write([]byte{'d'})

	fmt.Println("缓冲 IO")
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	bw.Write([]byte{'d'})
	if err := bw.Flush(); err != nil {
		log.Fatal(err)
	}
}

//无缓冲IO
//1
//1
//1
//1
//缓冲 IO
//3
//1
