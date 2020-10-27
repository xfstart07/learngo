package main

import (
	"fmt"
)

func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		4: func() int { return 10 },
	}
	fmt.Println(mf)
}
