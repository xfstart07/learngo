package main

import (
	"learngo/pkg/timeago/timeago"
	"time"

	"fmt"
)

func main() {
	t := time.Now().Add(42 * time.Second)

	now := time.Now()

	fmt.Println(now.Sub(t))

	s := timeago.English.Format(t)

	fmt.Println(s)
}
