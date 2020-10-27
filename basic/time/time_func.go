package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	now2 := now

	fmt.Println("now", now.After(now2))

	before := now.Add(-1 * time.Second)
	fmt.Println("before", now.After(before))

	after := now.Add(1 * time.Second)
	fmt.Println("after", now.After(after))

}
