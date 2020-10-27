package main

import "fmt"

func main() {
	multi(1, 2, 3)
}

func multi(options ...interface{}) {
	fmt.Println(options[1])
}
