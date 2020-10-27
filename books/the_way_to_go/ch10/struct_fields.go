package main

import (
	"fmt"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.1
	ms.str = "Chris"

	fmt.Println(ms)
}
