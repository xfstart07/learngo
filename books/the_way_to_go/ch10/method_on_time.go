package main

import "time"
import "fmt"

type myTime struct {
	time.Time // anonymous field
}

func main() {
	m := myTime{time.Now()}

	fmt.Println("Full time now:", m.String())

	fmt.Println("First 3 chars:", m.first3Chars())
}

func (_ myTime) first3Chars() string {
	return "11"
	// return t.Time.String()[0:3]
}

// Full time now: 2018-01-01 21:11:39.932549757 +0800 CST m=+0.000270328
// First 3 chars: 201
