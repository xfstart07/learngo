package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")

	go longWait()
	go shortWait()

	time.Sleep(10 * time.Second)
	fmt.Println("end main()")
}

func longWait() {
	fmt.Println("begin longwait()")
	time.Sleep(5 * 1e9)
	fmt.Println("end longwait()")
}

func shortWait() {
	fmt.Println("begin shortwait()")
	time.Sleep(2 * 1e9)
	fmt.Println("end shortwait()")
}
