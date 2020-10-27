package main

import (
	"fmt"
)

func main() {
	g := func(i int) {
		fmt.Println(i)
	}
	for i := 0; i < 4; i++ {
		g(i)
	}
}
