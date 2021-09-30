package main

import "fmt"

func main() {
	arr := []int{}

	for idx := range arr {
		fmt.Println(idx)
	}

	arrNil := []int{}
	arrNil = nil

	for idx := range arrNil {
		fmt.Println(idx)
	}

}
