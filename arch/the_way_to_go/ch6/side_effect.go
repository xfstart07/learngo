package main

import (
	"fmt"
)

func main() {
	n := 0
	reply := &n
	Multiply(1, 2, reply)
	fmt.Println(*reply)
}

func Multiply(a, b int, reply *int) {
	*reply = a + b
}
